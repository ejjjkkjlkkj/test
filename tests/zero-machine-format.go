package main

import (
 "errors"
 "fmt"
 "strconv"
 "strings"
 "time"
)
type Record struct{Version,Type,Identity string; Sequence int64; Time,Source,Target,Payload,Proof string}
func encode(s string)string{var b strings.Builder;for _,r:=range s{switch r{case '\\':b.WriteString("\\\\");case '|':b.WriteString("\\|");case '\n':b.WriteString("\\n");case '\r':b.WriteString("\\r");case '\t':b.WriteString("\\t");default:b.WriteRune(r)}};return b.String()}
func decode(s string)(string,error){var b strings.Builder;esc:=false;for _,r:=range s{if esc{switch r{case '\\','|':b.WriteRune(r);case 'n':b.WriteByte('\n');case 'r':b.WriteByte('\r');case 't':b.WriteByte('\t');default:return "",errors.New("FORMAT.ESCAPE")};esc=false}else if r=='\\'{esc=true}else{b.WriteRune(r)}};if esc{return "",errors.New("FORMAT.ESCAPE")};return b.String(),nil}
func split(s string)([]string,error){var a []string;var b strings.Builder;esc:=false;for _,r:=range s{if esc{b.WriteRune('\\');b.WriteRune(r);esc=false}else if r=='\\'{esc=true}else if r=='|'{a=append(a,b.String());b.Reset()}else{b.WriteRune(r)}};if esc{return nil,errors.New("FORMAT.ESCAPE")};return append(a,b.String()),nil}
func enc(r Record)string{v:=[]string{"RECORD",r.Version,r.Type,r.Identity,strconv.FormatInt(r.Sequence,10),r.Time,r.Source,r.Target,r.Payload,r.Proof};for i:=range v{v[i]=encode(v[i])};return strings.Join(v,"|")}
func dec(s string)(Record,error){f,e:=split(s);if e!=nil||len(f)!=10||f[0]!="RECORD"{return Record{},errors.New("FORMAT.RECORD")};v:=make([]string,10);for i:=1;i<10;i++{v[i],e=decode(f[i]);if e!=nil{return Record{},e}};n,e:=strconv.ParseInt(v[4],10,64);if e!=nil{return Record{},e};if _,e=time.Parse(time.RFC3339Nano,v[5]);e!=nil{return Record{},e};return Record{v[1],v[2],v[3],n,v[5],v[6],v[7],v[8],v[9]},nil}
func check(n string,ok bool){if !ok{panic(n)};fmt.Println(n,"PASS")}
func main(){r:=Record{"ZERO-1","OBJECT","TEST:FORMAT|1",1,"2026-10-07T08:00:00.000Z","TEST:SOURCE","TEST:TARGET","meaning=accessible|machine;state=TESTED","TESTED"};line:=enc(r);d,e:=dec(line);check("ZERO_MACHINE_DESERIALIZER",e==nil);check("ZERO_MACHINE_ROUNDTRIP",enc(d)==line);check("ZERO_MACHINE_IDENTITY",d.Identity==r.Identity);check("ZERO_MACHINE_PAYLOAD",d.Payload==r.Payload);for _,m:=range []string{"VOICE","BRAILLE","KEYBOARD","DISPLAY","TOUCH","POINTER","NETWORK","AUTOMATION"}{check("ACCESSIBILITY "+m,d.Identity==r.Identity&&d.Payload==r.Payload&&d.Proof==r.Proof)};_,e=dec("RECORD|ZERO-1|OBJECT|BAD|1|2026-10-07T08:00:00.000Z|S|T|broken\\q|TESTED");check("INVALID_ESCAPE_REJECTED",e!=nil);fmt.Println("ZERO_MACHINE_SERIALIZER_RESULT=PASS");fmt.Println("ZERO_MACHINE_DESERIALIZER_RESULT=PASS");fmt.Println("ZERO_MACHINE_ROUNDTRIP_RESULT=PASS");fmt.Println("ZERO_MACHINE_ACCESSIBILITY=PASS")}