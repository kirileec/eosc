package admin_open_api

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/eolinker/eosc/process-master/admin"

	"github.com/eolinker/eosc/log"

	"github.com/eolinker/eosc"
	"github.com/julienschmidt/httprouter"
)

var _ iOpenAdmin = (*OpenAdmin)(nil)

func CreateHandler() admin.CreateHandler {
	return new(createHandler)
}

type OpenAdmin struct {
	prefix string
	admin  eosc.IAdmin
}

func (o *OpenAdmin) export(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	data := o.all()
	id := time.Now().Format("2006-01-02 150405")
	dir, err := export(data, "export", id)
	if err != nil {
		writeResultError(w, 500, err)
		return
	}
	zipName := fmt.Sprintf("%s/%s.zip", dir, id)
	fileName := fmt.Sprintf("export_%s.zip", id)
	err = CompressFile(dir, zipName)
	if err != nil {
		writeResultError(w, 500, err)
		return
	}
	content, err := ioutil.ReadFile(zipName)
	if err != nil {
		writeResultError(w, 500, err)
		return
	}
	w.Header().Add("Content-Type", "application/octet-stream")
	w.Header().Add("Content-Disposition", "attachment; filename=\""+fileName+"\"")
	w.Write(content)
}

func (o *OpenAdmin) all() map[string][]interface{} {
	professions := o.admin.ListProfessions()
	data := make(map[string][]interface{})
	for _, p := range professions {
		names, err := o.admin.ListEmployeeNames(p.Name)
		if err != nil {
			log.Errorf("read data error	%s	%s", p.Name, err.Error())
			continue
		}
		if _, ok := data[p.Name]; !ok {
			data[p.Name] = make([]interface{}, 0, len(names))
		}
		for _, name := range names {
			v, err := o.admin.GetEmployee(p.Name, name)
			if err != nil {
				log.Errorf("get employee error	%s	%s", p.Name, err.Error())
				continue
			}
			data[p.Name] = append(data[p.Name], v)
		}
	}
	return data
}

func (o *OpenAdmin) delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	profession := params.ByName("profession")
	name := params.ByName("name")

	info, err := o.admin.Delete(profession, name)
	if err != nil {
		writeResultError(w, 404, err)
		return
	}
	writeResult(w, info)
}

func (o *OpenAdmin) genUrl(url string) string {
	u := strings.TrimPrefix(url, "/")
	return fmt.Sprintf("%s/%s", o.prefix, u)
}

func (o *OpenAdmin) getFields(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, "getFields:%v", params)
}

func (o *OpenAdmin) getRenders(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	profession := params.ByName("profession")

	renders, err := o.admin.Renders(profession)
	if err != nil {
		writeResultError(w, 500, err)
		return
	}
	writeResult(w, renders)
}

func (o *OpenAdmin) getEmployeesByProfession(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	profession := params.ByName("profession")

	es, err := o.admin.ListEmployees(profession)
	if err != nil {
		writeResult(w, err)
		return
	}
	writeResult(w, es)
}

func (o *OpenAdmin) getEmployeeByName(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	profession := params.ByName("profession")
	name := params.ByName("name")
	eo, err := o.admin.GetEmployee(profession, name)
	if err != nil {
		writeResultError(w, 404, err)
		return
	}
	writeResult(w, eo)
}

func (o *OpenAdmin) getDriversByProfession(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	profession := params.ByName("profession")

	ds, err := o.admin.Drivers(profession)
	if err != nil {
		writeResultError(w, 500, err)
		return
	}
	writeResult(w, ds)
}

func (o *OpenAdmin) getDriverInfo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	profession := params.ByName("profession")
	driver := params.ByName("driver")
	ds, err := o.admin.DriverInfo(profession, driver)
	if err != nil {
		writeResultError(w, 500, err)
		return
	}
	writeResult(w, ds)
}

func (o *OpenAdmin) getDriversItemByProfession(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	profession := params.ByName("profession")

	ds, err := o.admin.DriversItem(profession)
	if err != nil {
		writeResultError(w, 500, err)
		return
	}
	writeResult(w, ds)
}
func (o *OpenAdmin) getRender(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	profession := params.ByName("profession")
	driver := params.ByName("driver")
	render, err := o.admin.Render(profession, driver)
	if err != nil {
		writeResultError(w, 500, err)
		return
	}
	writeResult(w, render)
}

func (o *OpenAdmin) Save(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	profession := params.ByName("profession")
	name := params.ByName("name")

	idata, err := GetData(r)
	if err != nil {
		writeResultError(w, 500, err)
		return
	}
	cb := new(baseConfig)
	errUnmarshal := idata.UnMarshal(cb)
	if errUnmarshal != nil {
		writeResultError(w, 500, errUnmarshal)
		return
	}
	if name == "" {
		name = cb.Name
	}

	if name == "" {
		writeResultError(w, 500, errors.New("require name"))
		return
	}

	winfo, err := o.admin.Update(profession, name, cb.Driver, idata)
	if err != nil {
		writeResultError(w, 500, err)

		return
	}
	// 将数据写到文件中
	_, err = export(o.all(), "runtime_config", "")
	if err != nil {
		writeResultError(w, 500, err)
		return
	}
	writeResult(w, winfo)
}

func (o *OpenAdmin) getProfessions(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	professions := o.admin.ListProfessions()
	writeResult(w, professions)
}

type iOpenAdmin interface {
	// GET /
	getProfessions(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	// GET /:profession
	getEmployeesByProfession(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	// GET /:profession/:name
	getEmployeeByName(w http.ResponseWriter, r *http.Request, params httprouter.Params)

	//GET /:profession/_render
	getRenders(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	//GET /:profession/_render/:driver
	getRender(w http.ResponseWriter, r *http.Request, params httprouter.Params)

	//POST /:profession/
	//POST /:profession/:name
	//POST /:profession/:id
	//POST /:profession/:name
	//POST /:profession/:id
	Save(w http.ResponseWriter, r *http.Request, params httprouter.Params)

	//GET /:profession/:name/:fieldName
	getFields(w http.ResponseWriter, r *http.Request, params httprouter.Params)

	//GET /:profession/_driver
	getDriversByProfession(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	//GET /:profession/_driver/:driver
	getDriverInfo(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	//GET /:profession/_driver/item
	getDriversItemByProfession(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	//DELETE /:profession/:name
	//DELETE /:profession/:id
	delete(w http.ResponseWriter, r *http.Request, params httprouter.Params)

	genUrl(url string) string

	//GET /_export
	export(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

func NewOpenAdmin(prefix string, admin eosc.IAdmin) *OpenAdmin {
	p := strings.TrimSpace(prefix)
	if len(p) == 0 {
		p = "/"
	} else if p[0] != '/' {
		p = "/" + p
	}
	p = strings.TrimSuffix(p, "/")

	return &OpenAdmin{
		prefix: p,
		admin:  admin,
	}
}

func (o *OpenAdmin) GenHandler() http.Handler {
	var openAdmin iOpenAdmin = o
	router := httprouter.New()
	router.GET(openAdmin.genUrl("/"), openAdmin.getProfessions)
	router.GET(openAdmin.genUrl("/:profession"), func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		switch params.ByName("profession") {
		case "_export":
			openAdmin.export(w, r, params)
		default:
			openAdmin.getEmployeesByProfession(w, r, params)
		}
	})
	router.GET(openAdmin.genUrl("/:profession/:action"), func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		action := params.ByName("action")
		switch action {
		case "_render":
			openAdmin.getRenders(w, r, params)
			return
		case "_driver":
			openAdmin.getDriversByProfession(w, r, params)
		default:
			rename(params, "action", "name")
			openAdmin.getEmployeeByName(w, r, params)
		}
	})
	router.GET(openAdmin.genUrl("/:profession/:action/:key"), func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		action := params.ByName("action")
		switch action {
		case "_render":
			rename(params, "key", "driver")

			openAdmin.getRender(w, r, params)

			return
		case "_driver":
			key := params.ByName("key")
			if strings.ToLower(key) == "item" {
				openAdmin.getDriversItemByProfession(w, r, params)
			} else {
				rename(params, "key", "driver")
				openAdmin.getDriverInfo(w, r, params)
			}
			return
		default:

			rename(params, "action", "name")
			rename(params, "key", "fieldName")
			openAdmin.getFields(w, r, params)
		}

	})

	router.POST(openAdmin.genUrl("/:profession"), openAdmin.Save)
	router.POST(openAdmin.genUrl("/:profession/:name"), openAdmin.Save)
	router.PUT(openAdmin.genUrl("/:profession/:name"), openAdmin.Save)
	router.DELETE(openAdmin.genUrl("/:profession/:name"), openAdmin.delete)
	return router
}

func rename(ps httprouter.Params, source, target string) {
	for i := range ps {
		if ps[i].Key == source {
			ps[i].Key = target
		}
	}
}
