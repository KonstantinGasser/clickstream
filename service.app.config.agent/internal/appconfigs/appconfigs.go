package appconfigs

import (
	"context"
	"fmt"
)

const (
	FlagFunnel   = "funnel"
	FlagCampaign = "campaign"
	FlagBtnTime  = "btntime"
)

var (
	ErrInvalidFlag       = fmt.Errorf("provided config-flag is invalid")
	ErrNoReadWriteAccess = fmt.Errorf("user read/write access for AppToken")
	ErrNoReadAccess      = fmt.Errorf("user has no read access for AppToken")
)

type AppconfigRepo interface {
	Initialize(ctx context.Context, appConfig AppConfig) error
	GetById(ctx context.Context, uuid string, result interface{}) error
	UpdateByFlag(ctx context.Context, uuid, flag string, data interface{}) error
}

type AppConfig struct {
	AppUuid     string   `bson:"_id"`
	ConfigOwner string   `bson:"config_owner"`
	Funnel      []Stage  `bson:"funnel"`
	Campaign    []Record `bson:"campaign"`
	BtnTime     []BtnDef `bson:"btn_time"`
}

type Stage struct {
	Id         int32  `bson:"id"`
	Name       string `bson:"name"`
	Transition string `bson:"transition"`
}

type Record struct {
	Id     int32  `bson:"id"`
	Name   string `bson:"name"`
	Suffix string `bson:"suffix"`
}

type BtnDef struct {
	Id      int32  `bson:"id"`
	Name    string `bson:"name"`
	BtnName string `bson:"btn_name"`
}

// NewDefault creates a new AppConfig with only meta data
func NewDefault(appRefUuid, configOwner string) *AppConfig {
	return &AppConfig{
		AppUuid:     appRefUuid,
		ConfigOwner: configOwner,
		Funnel:      make([]Stage, 0),
		Campaign:    make([]Record, 0),
		BtnTime:     make([]BtnDef, 0),
	}
}

func (appConf *AppConfig) ApplyFunnel(stages ...Stage) {
	appConf.Funnel = stages
}

func (appConf *AppConfig) ApplyCampaign(records ...Record) {
	appConf.Campaign = records
}

func (appConf *AppConfig) ApplyBtnTime(btnDefs ...BtnDef) {
	appConf.BtnTime = btnDefs
}

// HasReadWrite checks if the provided user uuid is listed as owner of
// AppToken
func (appConfig AppConfig) HasReadWrite(userUuid string) error {
	if appConfig.ConfigOwner != userUuid {
		return ErrNoReadWriteAccess
	}
	return nil
}

// HasRead checks if the user has read access on the AppToken
func (appConfig AppConfig) HasRead(readWriteUuids ...string) error {
	for _, uuid := range readWriteUuids {
		if uuid == appConfig.AppUuid {
			return nil
		}
	}
	return ErrNoReadAccess
}

// HasReadOrWrite checks if the user has either read or write acces on the AppToken
func (appConfig AppConfig) HasReadOrWrite(userUuid string, readWriteUuids ...string) error {
	var err error
	if err = appConfig.HasRead(readWriteUuids...); err != nil {
		return err
	}

	if err = appConfig.HasReadWrite(userUuid); err != nil {
		return err
	}
	return nil
}