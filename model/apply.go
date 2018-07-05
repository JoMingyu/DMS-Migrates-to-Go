package model

import "DMS-Migrates-to-Go/db"

// ExtensionApplyModel 구조체는 ExtensionApply11Model과 Extension12Model을 위한 기반 구조체입니다.
type ExtensionApplyModel struct {
	Student StudentModel
	Class   int
	Seat    int
}

// StayApplyModel 구조체는 기숙사 잔류 신청 정보를 관리합니다.
type StayApplyModel struct {
	Student StudentModel
	Value   int
}

// GoingoutApplyModel 구조체는 매 주말마다의 외출 신청 정보를 관리합니다.
type GoingoutApplyModel struct {
	Student    StudentModel
	OnSaturday bool
	OnSunday   bool
}

var (
	// ExtensionApply11Col 은 ExtensionApply11Model에 대한 collection을 참조합니다.
	ExtensionApply11Col = db.DB().C("extension_apply_11")

	// ExtensionApply12Col 은 ExtensionApply12Model에 대한 collection을 참조합니다.
	ExtensionApply12Col = db.DB().C("extension_apply_12")

	// StayApplyCol 은 StayApplyModel에 대한 collection을 참조합니다.
	StayApplyCol = db.DB().C("stay_apply")

	// GoingoutApplyCol 은 GoingoutApplyModel에 대한 collection을 참조합니다.
	GoingoutApplyCol = db.DB().C("goingout_apply")
)
