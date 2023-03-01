package storefront

import (
	"errors"

	"github.com/MyriadFlow/storefront-gateway/config/smartcontract"
	"github.com/MyriadFlow/storefront-gateway/util/pkg/logwrapper"
)

type tRole int

var (
	ErrRoleNotExist = errors.New("role does not exist")
)

const (
	CREATOR_ROLE  tRole = iota
	ADMIN_ROLE    tRole = iota
	OPERATOR_ROLE tRole = iota
)

type tRoles map[tRole][32]byte

var roles tRoles = tRoles{}
var initiated = false

func GetRole(role tRole) ([32]byte, error) {
	if !initiated {
		InitRolesId()
	}
	v, ok := roles[role]
	if !ok {
		return [32]byte{}, ErrRoleNotExist
	}
	return v, nil
}
func InitRolesId() {
	client, err := smartcontract.GetClient()
	if err != nil {
		logwrapper.Fatalf("failed to client, error: %v", err.Error())
	}
	instance, err := GetInstance(client)
	if err != nil {
		logwrapper.Fatalf("failed to get instance for %v , error: %v", "STOREFRONT", err.Error())
	}
	creatorRoleId, err := instance.STOREFRONTCREATORROLE(nil)
	if err != nil {
		logwrapper.Fatalf("Failed to get %v, error: %v", "STOREFRONTCREATORROLE", err.Error())
	}
	roles[CREATOR_ROLE] = creatorRoleId
	adminRoleId, err := instance.STOREFRONTADMINROLE(nil)
	if err != nil {
		logwrapper.Fatalf("Failed to get %v, error: %v", "STOREFRONTADMINROLE", err.Error())
	}
	roles[ADMIN_ROLE] = adminRoleId

	operatorRoleId, err := instance.STOREFRONTOPERATORROLE(nil)
	if err != nil {
		logwrapper.Fatalf("Failed to get %v, error: %v", "STOREFRONTOPERATORROLE", err.Error())
	}
	roles[OPERATOR_ROLE] = operatorRoleId
}
