package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

/*
	Паттерн определяет интерфейс создания структуры, реализующей некоторый интерфейс, и при этом оставляет возможность
	выбора на основании какой структуры создавать экземпляр
	Т.е. делает код создания объектов более универсальным без привязки к конкретным структурам.
*/

// Product
type Computer interface {
	CPU() string
	RAM() string
	HDD() string
}

// Concrete products
type PC struct {
	cpu      string
	ram      string
	hdd      string
	monitor  string
	mouse    string
	keyboard string
}

func (pc *PC) CPU() string {
	return pc.cpu
}

func (pc *PC) RAM() string {
	return pc.ram
}

func (pc *PC) HDD() string {
	return pc.hdd
}

type Tablet struct {
	cpu     string
	ram     string
	hdd     string
	stylus  string
	display string
}

func (t *Tablet) CPU() string {
	return t.cpu
}

func (t *Tablet) RAM() string {
	return t.ram
}

func (t *Tablet) HDD() string {
	return t.hdd
}

type Smartphone struct {
	cpu     string
	ram     string
	hdd     string
	camera  string
	display string
}

func (s *Smartphone) CPU() string {
	return s.cpu
}

func (s *Smartphone) RAM() string {
	return s.ram
}

func (s *Smartphone) HDD() string {
	return s.hdd
}

// Factory
type ComputerFactory interface {
	NewComputer() Computer
}

// Concrete factories
type PCFactory struct {
	cpu      string
	ram      string
	hdd      string
	monitor  string
	mouse    string
	keyboard string
}

func NewPCFactory(cpu string, ram string, hdd string, monitor string, mouse string, keyboard string) *PCFactory {
	return &PCFactory{
		cpu:      cpu,
		ram:      ram,
		hdd:      hdd,
		monitor:  monitor,
		mouse:    mouse,
		keyboard: keyboard,
	}
}

func (pcf *PCFactory) NewComputer() Computer {
	return &PC{
		cpu:      pcf.cpu,
		ram:      pcf.ram,
		hdd:      pcf.hdd,
		monitor:  pcf.monitor,
		mouse:    pcf.mouse,
		keyboard: pcf.keyboard,
	}
}

type TabletFactory struct {
	cpu     string
	ram     string
	hdd     string
	stylus  string
	display string
}

func NewTabletFactory(cpu string, ram string, hdd string, stylus string, display string) *TabletFactory {
	return &TabletFactory{
		cpu:     cpu,
		ram:     ram,
		hdd:     hdd,
		stylus:  stylus,
		display: display,
	}
}

func (tf *TabletFactory) NewComputer() Computer {
	return &Tablet{
		cpu:     tf.cpu,
		ram:     tf.ram,
		hdd:     tf.hdd,
		stylus:  tf.stylus,
		display: tf.stylus,
	}
}

type SmartphoneFactory struct {
	cpu     string
	ram     string
	hdd     string
	camera  string
	display string
}

func NewSmartphoneFactory(cpu string, ram string, hdd string, camera string, display string) *SmartphoneFactory {
	return &SmartphoneFactory{
		cpu:     cpu,
		ram:     ram,
		hdd:     hdd,
		camera:  camera,
		display: display,
	}
}

func (sp *SmartphoneFactory) NewComputer() Computer {
	return &Smartphone{
		cpu:     sp.cpu,
		ram:     sp.ram,
		hdd:     sp.hdd,
		camera:  sp.camera,
		display: sp.display,
	}
}

func main() {
	factories := []ComputerFactory{
		NewPCFactory("I7", "HyperX DDR5", "Toshiba 2TB", "LG UltraSync", "Logitech G Pro", "Varmilo SK87"),
		NewTabletFactory("Snapdragon 855", "Samsung DDR4", "512 GB", "Apple Pen", "IPS 1920x1080"),
		NewSmartphoneFactory("Apple Bionic A20", "Samsung DDR5", "256 GB", "Leica", "OLED 2160x1440"),
	}

	for _, factory := range factories {
		computer := factory.NewComputer()
		fmt.Println(computer, computer.CPU(), computer.HDD(), computer.RAM())
	}
}
