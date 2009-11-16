package complex

const Msg1="Hello world"

type Msg struct { msg string;}

func (m *Msg) GetMsg() string { return m.msg; }

func (m *Msg) SetMsg(value string) { 
	if value == "" { m.msg = value; } else { m.msg = "empty"; }
}

