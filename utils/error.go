package utils

type ErrorResponse struct {
	Err string
}

type error interface {
	Error() string
}



func GetError(code int, lang string) (map[string]interface{}, map[string]interface{}){
	errors := map[int]map[string]string{
        200: map[string]string{
        	"fr": "OK",
        	"en": "OK",
        },
        201: map[string]string{
        	"fr": "Crée",
        	"en": "Created",
        },
        400: map[string]string{
        	"fr": "Requete Invalide",
        	"en": "Bad Request",
        },
        401: map[string]string{
        	"fr": "Non authorizé",
        	"en": "Unauthorized",
        },
        403: map[string]string{
        	"fr": "Accès restreint",
        	"en": "Forbidden",
        },
        404: map[string]string{
        	"fr": "Aucun résultat",
        	"en": "Not found",
        },
        500: map[string]string{
        	"fr": "Erreur Serveur",
        	"en": "Internal Server Error",
        },

    }


	if _, ok:= errors[code]; !ok {
		var resp = map[string]interface{}{"status": 500, "message": errors[500]["en"]}
		return nil, resp
	}

	if _, ok2:= errors[code][lang]; !ok2 {
		var resp = map[string]interface{}{"status": 500, "message": errors[500]["en"]}
		return nil, resp
	}

	var resp = map[string]interface{}{"status": code, "message": errors[code][lang]}
	return resp, nil
}

