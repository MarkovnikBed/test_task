package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"medods/internal/tokens"
)

func (h *Handler) Refresh(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// вытягиваем токены из кук
	refreshToken, err := r.Cookie("RT")
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(Response{Err: "не получилось получить refresh токен"})
		return
	}
	accessToken, err := r.Cookie("AT")
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(Response{Err: "не получилось получить access токен"})
		return
	}
	// проверяем наличие токена из кук в БД
	issueAt, id, ip, err := h.rep.ValidRefreshToken(refreshToken.Value)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(Response{Err: err.Error()})
		return
	}
	// проверяем создан ли был AT вместе с RT
	ipAT, err := tokens.CompareRTandAT(accessToken.Value, id, issueAt)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(Response{Err: err.Error()})
		return
	}
	// пересоздаём AT
	newAT, err := tokens.GetNewAccessToken(ip, id, issueAt)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(Response{Err: "не удалось пересоздать AT"})
		return
	}
	// добавляем новый AT в куки
	http.SetCookie(w, &http.Cookie{
		Name:     "AT",
		Value:    newAT,
		Path:     "/",
		Expires:  time.Now().Add(15 * time.Minute),
		Secure:   true,
		HttpOnly: true,
	})
	if ipAT != ip {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Response{Warning: "На ваш email отправлено предупреждение"})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Response{Success: "access токен обновлён"})
	}

}
