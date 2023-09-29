package main

import "testing"

type testCase struct {
	Alphabet         string
	Message          string
	Key              string
	EncryptedMessage string
}

func Test_EncryptVigenere(t *testing.T) {
	testCases := []testCase{
		{
			"АБВГДЕЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ ",
			"В КНИГАХ МНОГО ТАЙН И РАЗГАДОК",
			"ЛЕВ",
			"НДМШНЕЛЪБЧТРОУБЭЕЛШДККХВТИВПУМ",
		},
		{
			"АБВГДЕЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ ",
			"ХРАНЯТ МНОГО ЭКСПОНАТОВ",
			"ЕЛИСЕЙ",
			"ЪЫИЮГЫДЧХЯИЧДЗТБФЧТЛЪЯЗ",
		},
		{
			"АбвгдеЖЗИЙКлМНОПРСТУФХЦЧШЩЪЫЬЭюя ",
			"хранят МНОГО эКСПОНАТОв",
			"Елисей",
			"ЪЫИЮГЫДЧХЯИЧДЗТБФЧТЛЪЯЗ",
		},
	}

	for _, testCase := range testCases {
		encryptedMessage, err := EncryptVigenere(testCase.Alphabet, testCase.Message, testCase.Key)
		if err != nil {
			t.Errorf("An error during test case '%v' occured: %s", testCase, err)
		}

		if encryptedMessage != testCase.EncryptedMessage {
			t.Errorf("Expected '%s' encrypted message got '%s' for test case '%v'", testCase.EncryptedMessage, encryptedMessage, testCase)
		}
	}
}
