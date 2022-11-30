package main

import (
	"math"
	"testing"
)

func TestInsertMapVoteDB(t *testing.T) {
	m := make(MapVoteDB)
	v := Vote{"", 1}
	v, _ = m.Insert(v)
	if v.Choice != 1 {
		t.Errorf("got choice %d want %d\n", v.Choice, 1)
	}
	if v.Id == "" || len(v.Id) != 36 {
		t.Errorf("%q is an invalid id\n", v.Id)
	}
}

func TestGetMapVoteDB(t *testing.T) {
	m := make(MapVoteDB)
	want, _ := m.Insert(Vote{"", 1})
	got, _ := m.Get(want.Id)
	if got.Id != want.Id {
		t.Errorf("got vote id %q want %q\n", got.Id, want.Id)
	}
	if got.Choice != want.Choice {
		t.Errorf("got vote choice %d want %d\n", got.Choice, want.Choice)
	}
}

func TestGetEmptyMapVoteDB(t *testing.T) {
	m := make(MapVoteDB)
	_, err := m.Get("")
	if err == nil {
		t.Logf("expected error\n")
	}
}

func TestGetAllMapVoteDB(t *testing.T) {
	m := make(MapVoteDB)
	m.Insert(Vote{"", 1})
	m.Insert(Vote{"", 2})
	m.Insert(Vote{"", 1})
	votes, _ := m.GetAll()
	if len(votes) != 3 {
		t.Errorf("got slice size %d want %d\n", len(votes), 3)
	}
}

func TestGetAllEmptyMapVoteDB(t *testing.T) {
	m := make(MapVoteDB)
	votes, _ := m.GetAll()
	if len(votes) != 0 {
		t.Errorf("got slice size %d want %d\n", len(votes), 0)
	}
}

func TestGetPercentagesMapVoteDB(t *testing.T) {
	m := make(MapVoteDB)
	m.Insert(Vote{"", 1})
	m.Insert(Vote{"", 2})
	m.Insert(Vote{"", 1})
	percentages, _ := m.GetPercentages()
	if math.Abs(percentages[0].Percentage-66.66) <= 1e-9 {
		t.Errorf("got choice `1` with percentage %f want %f\n", percentages[0].Percentage, 66.66)
	}
	if math.Abs(percentages[1].Percentage-33.33) <= 1e-9 {
		t.Errorf("got choice `2` with percentage %f want %f\n", percentages[1].Percentage, 33.33)
	}
}

func TestGetPercentagesEmptyMapVoteDB(t *testing.T) {
	m := make(MapVoteDB)
	percentages, _ := m.GetPercentages()
	if len(percentages) != 0 {
		t.Errorf("got percentage list size %d want %d\n", len(percentages), 0)
	}
}
