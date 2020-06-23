package builder
type People interface {
	SetA(s string)
	SetB(s string)
	SetC(s string)
}
type Man struct {
	A string
	B string
	C string
}

func (m *Man)SetA(s string)  {
	m.A=s
}

func (m *Man)SetB(s string)  {
	m.B=s
}

func (m *Man)SetC(s string)  {
	m.C=s
}
type Builder interface {
	Init()
	BuildA()
	BuildB()
	BuildC()
	GetPeople() People
}
type ManBuilder struct {
	P People
}

func (m *ManBuilder)Init()  {
	m.P=new(Man)
}
func (m *ManBuilder)BuildA()  {
	m.P.SetA("this is a man")
}
func (m *ManBuilder)BuildB()  {
	m.P.SetB("this is a man")
}
func (m *ManBuilder)BuildC()  {
	m.P.SetC("this is a man")
}
func (m *ManBuilder) GetPeople() People {
	return m.P
}
type Director struct {
	builder Builder
}

func (d *Director)SetBuilder(b Builder)  {
	d.builder=b
}
func (d *Director)Construct()People  {
	d.builder.Init()
	d.builder.BuildA()
	d.builder.BuildB()
	d.builder.BuildC()
	return d.builder.GetPeople()
}

