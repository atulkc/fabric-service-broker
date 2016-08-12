package bosh_test

import (
	"testing"

	"github.com/atulkc/fabric-service-broker/bosh"

	. "gopkg.in/go-playground/assert.v1"
)

const (
	boshStemcell = "mystemcell"
	boshUuid     = "uuid-1"
	vmType       = "vmtype"
	networkNames = "net1,net2,net3"
	directorUrl  = "http://the-bosh-director"
)

func TestNewDetails(t *testing.T) {
	boshDetails := bosh.NewDetails(boshStemcell, boshUuid, vmType, networkNames, directorUrl)
	NotEqual(t, boshDetails, nil)
	Equal(t, boshDetails.StemcellName, boshStemcell)
	Equal(t, boshDetails.DirectorUUID, boshUuid)
	Equal(t, boshDetails.Vmtype, vmType)
	Equal(t, len(boshDetails.NetworkNames), 3)
	err := boshDetails.Validate()
	Equal(t, err, nil)
}

func TestDetailsValidate_Stemcell(t *testing.T) {
	boshDetails := bosh.NewDetails("", boshUuid, vmType, networkNames, directorUrl)
	NotEqual(t, boshDetails, nil)
	err := boshDetails.Validate()
	NotEqual(t, err, nil)
	Equal(t, err.Error(), "StemcellName cannot be empty")
}

func TestDetailsValidate_UUID(t *testing.T) {
	boshDetails := bosh.NewDetails(boshStemcell, "", vmType, networkNames, directorUrl)
	NotEqual(t, boshDetails, nil)
	err := boshDetails.Validate()
	NotEqual(t, boshDetails, nil)
	Equal(t, err.Error(), "DirectorUUID cannot be empty")
}

func TestDetailsValidate_VmType(t *testing.T) {
	boshDetails := bosh.NewDetails(boshStemcell, boshUuid, "", networkNames, directorUrl)
	NotEqual(t, boshDetails, nil)
	err := boshDetails.Validate()
	NotEqual(t, boshDetails, nil)
	Equal(t, err.Error(), "Vmtype cannot be empty")
}

func TestDetailsValidate_NetworkNames(t *testing.T) {
	boshDetails := bosh.NewDetails(boshStemcell, boshUuid, vmType, "", directorUrl)
	NotEqual(t, boshDetails, nil)
	err := boshDetails.Validate()
	NotEqual(t, boshDetails, nil)
	Equal(t, err.Error(), "Invalid network name in the list")
}

func TestDetailsValidate_DirectorUrl(t *testing.T) {
	boshDetails := bosh.NewDetails(boshStemcell, boshUuid, vmType, networkNames, "")
	NotEqual(t, boshDetails, nil)
	err := boshDetails.Validate()
	NotEqual(t, boshDetails, nil)
	Equal(t, err.Error(), "BoshDirectorUrl cannot be empty")
}