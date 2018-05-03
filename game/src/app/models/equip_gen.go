/*
 * Generated by generate_tables
 * Warning: don't modify this file
 */

package models

import (
	. "app/consts"
	"fmt"
	"github.com/rs/xid"
	. "goslib/base_model"
	"gslib/player"
)

type EquipModel struct {
	Ctx  *player.Player
	Data *Equip
}

func FindEquip(ctx *player.Player, uuid string) *EquipModel {
	if model := ctx.Store.Get([]string{"models", "equips"}, uuid); model != nil {
		return model.(*EquipModel)
	}
	return nil
}

func CreateEquip(ctx *player.Player, data *Equip) *EquipModel {
	if data.Uuid == "" {
		data.Uuid = xid.New().String()
	}
	model := &EquipModel{
		Ctx:  ctx,
		Data: data,
	}
	ctx.Store.Set([]string{"models", "equips"}, data.Uuid, model)
	return model
}

func (self *EquipModel) GetUuid() string {
	return self.Data.Uuid
}

func (self *EquipModel) GetTableName() string {
	return "equips"
}

func (self *EquipModel) Save() {
	self.Ctx.Store.UpdateStatus("equips", self.GetUuid(), STATUS_UPDATE)
}

func (self *EquipModel) Delete() {
	self.Ctx.Store.Del([]string{"models", "equips"}, self.GetUuid())
	self.Ctx.Store.UpdateStatus("equips", self.GetUuid(), STATUS_DELETE)
}

func (self *EquipModel) SqlForRec(status int8) string {
	data := self.Data
	switch status {
	case STATUS_DELETE:
		return fmt.Sprintf("DELETE FROM `equips` WHERE `uuid`='%s'", data.Uuid)
	case STATUS_CREATE:
		return fmt.Sprintf("INSERT INTO `equips` (uuid, user_id, level, conf_id, evolves, equips, exp) VALUES ('%s', '%s', %d, %d, '%s', '%s', %d)", data.Uuid, data.UserId, data.Level, data.ConfId, data.Evolves, data.Equips, data.Exp)
	case STATUS_UPDATE:
		return fmt.Sprintf("UPDATE `equips` SET user_id='%s', level=%d, conf_id=%d, evolves='%s', equips='%s', exp=%d WHERE `uuid`='%s'", data.UserId, data.Level, data.ConfId, data.Evolves, data.Equips, data.Exp, data.Uuid)
	}
	return ""
}