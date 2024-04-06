package router

import (
	"GoLoanApp/constanta"
	"net/http"
)

func Middleware(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		CORSOriginHandler(&responseWriter)
		responseWriter.Header().Set("Content-Type", "application/json")
		responseWriter.Header().Set(constanta.RequestCacheControl, "no-cache")
		if request.Method == "OPTIONS" {
			return
		} else {
			//var contextModel *applicationModel.ContextModel
			defer func() {
				if r := recover(); r != nil {
					//util2.InputLog(errorMode.GenerateRecoverError(), contextModel.LoggerModel)
				}
			}()

			request.Header.Set(constanta.RequestCacheControl, "no-cache")
			//timestamp := time.Now()

			nextHandler.ServeHTTP(responseWriter, request)
		}
	})
}

func CORSOriginHandler(responseWriter *http.ResponseWriter) {
	(*responseWriter).Header().Set("Access-Control-Allow-Origin", "*")
	(*responseWriter).Header().Set("Access-Control-Allow-Headers", "origin, content-type, accept, authorization")
	(*responseWriter).Header().Set("Access-Control-Allow-Credentials", "true")
	(*responseWriter).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, HEAD")
	(*responseWriter).Header().Set("Access-Control-Max-Age", "1209600")
}
