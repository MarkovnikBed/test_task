package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"medods/internal/tokens"
)

func (h *Handler) Sign(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.FormValue("GUID")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&Response{Err: "не указан GUID пользователя"})
		return
	}
	// получаем RT
	refreshToken, err := tokens.GetRefreshToken(id, r.RemoteAddr)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(&Response{Err: err.Error()})
		return
	}
	//получаем AT
	accessToken, err := tokens.GetAccessToken(id, r.RemoteAddr)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(&Response{Err: err.Error()})
		return
	}
	// заносим полученный RT в БД в виде bcrypt-хэша
	err = h.rep.InsertRT(refreshToken, id)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(&Response{Err: "Не получилось сохранить токен в базе данных. " + err.Error()})
		return
	}
	// устанавливаем токены в куки
	http.SetCookie(w, &http.Cookie{
		Name:     "RT",
		Value:    refreshToken,
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour),
		Secure:   true,
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "AT",
		Value:    accessToken,
		Path:     "/",
		Expires:  time.Now().Add(15 * time.Minute),
		Secure:   true,
		HttpOnly: true,
	})

}
