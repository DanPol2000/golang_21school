package main

//#include <stdio.h>
//#include <stdlib.h>
//#include <string.h>
//
//unsigned int i;
//unsigned int argscharcount = 0;
//
//char *ask_cow(char phrase[]) {
//int phrase_len = strlen(phrase);
//char *buf = (char *)malloc(sizeof(char) * (160 + (phrase_len + 2) * 3));
//strcpy(buf, " ");
//
//for (i = 0; i < phrase_len + 2; ++i) {
//strcat(buf, "_");
//}
//
//strcat(buf, "\n< ");
//strcat(buf, phrase);
//strcat(buf, " ");
//strcat(buf, ">\n ");
//
//for (i = 0; i < phrase_len + 2; ++i) {
//strcat(buf, "-");
//}
//strcat(buf, "\n");
//strcat(buf, "        \\   ^__^\n");
//strcat(buf, "         \\  (oo)\\_______\n");
//strcat(buf, "            (__)\\       )\\/\\\n");
//strcat(buf, "                ||----w |\n");
//strcat(buf, "                ||     ||\n");
//return buf;
//}
import "C"
import (
	"bytes"
	"unsafe"
)

func ask_cow() string {
	phrase := C.CString("Thank you!")
	phraseLen := C.int(C.strlen(phrase))

	ptr := C.ask_cow(phrase)

	b := C.GoBytes(unsafe.Pointer(ptr), 160+(phraseLen+2)*3)
	b = bytes.Trim(b, "\x00")

	return string(b)
}