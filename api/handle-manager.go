package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/girithc/pronto-go/types"
)

func (s *Server) handleSendOtpManagerMSG91(res http.ResponseWriter, req *http.Request) error {
	new_req := new(types.Create_Customer)
	if err := json.NewDecoder(req.Body).Decode(new_req); err != nil {
		fmt.Println("Error in Decoding req.body in handleSendOtpMSG91")
		return err
	}

	result, err := s.store.SendOtpManagerMSG91(new_req.Phone)
	if err != nil {
		return err
	}

	return WriteJSON(res, http.StatusOK, result)
}

func (s *Server) handleVerifyOtpManagerMSG91(res http.ResponseWriter, req *http.Request) error {
	new_req := new(types.MobileOtp)
	if err := json.NewDecoder(req.Body).Decode(new_req); err != nil {
		fmt.Println("Error in Decoding req.body in handleVerifyOtpMSG91")
		return err
	}

	result, err := s.store.VerifyOtpManagerMSG91(new_req.Phone, new_req.Otp, new_req.FCM)
	if err != nil {
		return err
	}

	return WriteJSON(res, http.StatusOK, result)
}

func (s *Server) HandleManagerLogin(res http.ResponseWriter, req *http.Request) error {
	// Preprocessing

	new_req := new(types.ManagerFCM)

	if err := json.NewDecoder(req.Body).Decode(new_req); err != nil {
		fmt.Println("Error in Decoding req.body in HandleManagerLogin()")
		return err
	}

	packer, err := s.store.GetManagerByPhone(new_req.Phone, new_req.FCM)
	if err != nil {
		return err
	}

	// Check if User Exists

	return WriteJSON(res, http.StatusOK, packer)
}

func (s *Server) HandleManagerItems(res http.ResponseWriter, req *http.Request) error {
	items, err := s.store.GetManagerItems()
	if err != nil {
		return err
	}

	// Check if User Exists

	return WriteJSON(res, http.StatusOK, items)
}

func (s *Server) HandleManagerGetItem(res http.ResponseWriter, req *http.Request) error {
	new_req := new(GetItem)
	if err := json.NewDecoder(req.Body).Decode(new_req); err != nil {
		fmt.Println("Error in Decoding req.body in HandleManagerGetItem()")
		return err
	}

	item, err := s.store.GetManagerItem(new_req.ID)
	if err != nil {
		return err
	}

	return WriteJSON(res, http.StatusOK, item)
}

type GetItem struct {
	ID int `json:"id"`
}

func (s *Server) handleManagerInitShelfBasic(res http.ResponseWriter, req *http.Request) error {
	new_req := new(types.GetShelf)
	if err := json.NewDecoder(req.Body).Decode(new_req); err != nil {
		fmt.Println("Error in Decoding req.body in handleManagerInitShelf")
		return err
	}

	result, err := s.store.ManagerInitShelf(new_req.StoreId)
	if err != nil {
		return err
	}

	return WriteJSON(res, http.StatusOK, result)
}

func (s *Server) handleManagerAssignItemShelfBasic(res http.ResponseWriter, req *http.Request) error {
	new_req := new(types.AssignItemShelf)
	if err := json.NewDecoder(req.Body).Decode(new_req); err != nil {
		fmt.Println("Error in Decoding req.body in handleManagerAssignItemShelf")
		return err
	}

	result, err := s.store.ManagerAssignItemToShelf(*new_req)
	if err != nil {
		return err
	}

	return WriteJSON(res, http.StatusOK, result)
}
