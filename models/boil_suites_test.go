// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Persons", testPersons)
	t.Run("Things", testThings)
}

func TestDelete(t *testing.T) {
	t.Run("Persons", testPersonsDelete)
	t.Run("Things", testThingsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Persons", testPersonsQueryDeleteAll)
	t.Run("Things", testThingsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Persons", testPersonsSliceDeleteAll)
	t.Run("Things", testThingsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Persons", testPersonsExists)
	t.Run("Things", testThingsExists)
}

func TestFind(t *testing.T) {
	t.Run("Persons", testPersonsFind)
	t.Run("Things", testThingsFind)
}

func TestBind(t *testing.T) {
	t.Run("Persons", testPersonsBind)
	t.Run("Things", testThingsBind)
}

func TestOne(t *testing.T) {
	t.Run("Persons", testPersonsOne)
	t.Run("Things", testThingsOne)
}

func TestAll(t *testing.T) {
	t.Run("Persons", testPersonsAll)
	t.Run("Things", testThingsAll)
}

func TestCount(t *testing.T) {
	t.Run("Persons", testPersonsCount)
	t.Run("Things", testThingsCount)
}

func TestHooks(t *testing.T) {
	t.Run("Persons", testPersonsHooks)
	t.Run("Things", testThingsHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Persons", testPersonsInsert)
	t.Run("Persons", testPersonsInsertWhitelist)
	t.Run("Things", testThingsInsert)
	t.Run("Things", testThingsInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("PersonToThings", testPersonToManyThings)
	t.Run("ThingToPersons", testThingToManyPersons)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("PersonToThings", testPersonToManyAddOpThings)
	t.Run("ThingToPersons", testThingToManyAddOpPersons)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("PersonToThings", testPersonToManySetOpThings)
	t.Run("ThingToPersons", testThingToManySetOpPersons)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("PersonToThings", testPersonToManyRemoveOpThings)
	t.Run("ThingToPersons", testThingToManyRemoveOpPersons)
}

func TestReload(t *testing.T) {
	t.Run("Persons", testPersonsReload)
	t.Run("Things", testThingsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Persons", testPersonsReloadAll)
	t.Run("Things", testThingsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Persons", testPersonsSelect)
	t.Run("Things", testThingsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Persons", testPersonsUpdate)
	t.Run("Things", testThingsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Persons", testPersonsSliceUpdateAll)
	t.Run("Things", testThingsSliceUpdateAll)
}
