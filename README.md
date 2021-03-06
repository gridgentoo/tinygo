Оригинальный репозиторий.  
https://github.com/tinygo-org/tinygo

Tinygo - компилятор языка Go для областей, в которых необходимо компактное представление результирующего кода и низкое потребление ресурсов, таких как микроконтроллеры и компактные однопроцессорные системы. Код распространяется под лицензией BSD.

Компиляция для различных целевых платформ реализована при помощи LLVM, а для поддержки языка применяются библиотеки, применяемые в основном инструментарии от проекта Go. Скомпилированная программа напрямую может запускаться на микроконтроллерах, позволяя применять Go в качестве языка для написания сценариев автоматизации.

Мотивом создания нового проекта послужило желание использовать привычный для себя язык Go на компактных устройствах - разработчики рассудили, что если существует вариант Python для микроконтроллеров, то почему бы не создать подобное для языка Go. Go выбран вместо Rust так как он более прост в изучении, предоставляет независимую от реализаций потоков поддержку распараллеливания на основе сопрограмм и предлагает обширную стандартную библиотеку ("батарейки входят в комплект").

В текущем виде поддерживается 15 моделей микроконтроллеров, включая различные платы Adafruit, Arduino, BBC micro:bit, ST Micro, Digispark, Nordic Semiconductor, Makerdiary и Phytec. Программы также могут быть собраны для запуска в браузере в формате WebAssembly и в виде исполняемых файлов для Linux. Поддержки контроллеров ESP8266/ESP32 пока нет, но отдельно развивается проект по добавлении поддержки чипа Xtensa в LLVM, который пока отмечен как нестабильный и не готовых для интеграции с TinyGo.

Ключевые цели проекта:

Генерация очень компактных исполняемых файлов;
Поддержка наиболее распространённых моделей плат микроконтроллеров;
Возможность применения для Web;
Поддержка CGo с минимальными накладными расходами при вызове функций на языке Си;
Поддержка большей части стандартных пакетов и возможность компиляции типового существующего кода без его изменения.
Не входит в число основных целей поддержка многоядерных систем, эффективный запуск огромного числа сопрограмм (сам по себе запуск сопрограмм поддерживается в полной мере), достижение уровня производительности эталонного компилятора gc (оптимизация отдаётся на откуп LLVM и в некоторых применениях Tinygo может оказаться быстрее gc) и полная совместимость со всеми приложениями на Go.

Основным отличием от похожего компилятора emgo является попытка сохранения оригинальной модели управления памятью Go с использованием сборщика мусора и задействование LLVM для генерации эффективного кода вместо компиляции в представление на языке Си. Tinygo также предлагает новую runtime-библиотеку, в которой реализованы оптимизированные для компактных систем планировщик, система распределения памяти и обработчики строк. На базе нового runtime пересозданы некоторые пакеты, такие как sync и reflect.

Из изменений в выпуске 0.7 отмечается реализация команды "tinygo test", обеспечение поддержки сборки мусора для большинства целевых плат (на базе ARM Cortex-M) и WebAssembly, поддержка платы HiFive1 rev B на основе архитектуры RISC-V и платы Arduino nano33, улучшение поддержки языка (поддержка битовых полей с использованием геттеров и сеттеров, поддержка анонимных структур).

# TinyGo - Go compiler for small places

[![CircleCI](https://circleci.com/gh/tinygo-org/tinygo/tree/dev.svg?style=svg)](https://circleci.com/gh/tinygo-org/tinygo/tree/dev) [![Build Status](https://dev.azure.com/tinygo/tinygo/_apis/build/status/tinygo-CI?branchName=dev)](https://dev.azure.com/tinygo/tinygo/_build/latest?definitionId=1&branchName=dev)

TinyGo is a Go compiler intended for use in small places such as microcontrollers, WebAssembly (Wasm), and command-line tools.

It reuses libraries used by the [Go language tools](https://golang.org/pkg/go/) alongside [LLVM](http://llvm.org) to provide an alternative way to compile programs written in the Go programming language.

Here is an example program that blinks the built-in LED when run directly on any supported board with onboard LED:

```go
package main

import (
    "machine"
    "time"
)

func main() {
    led := machine.LED
    led.Configure(machine.PinConfig{Mode: machine.PinOutput})
    for {
        led.Low()
        time.Sleep(time.Millisecond * 1000)

        led.High()
        time.Sleep(time.Millisecond * 1000)
    }
}
```

The above program can be compiled and run without modification on an Arduino Uno, an Adafruit ItsyBitsy M0, or any of the supported boards that have a built-in LED, just by setting the correct TinyGo compiler target. For example, this compiles and flashes an Arduino Uno:

```shell
tinygo flash -target arduino examples/blinky1
```

## Installation

See the [getting started instructions](https://tinygo.org/getting-started/) for information on how to install TinyGo, as well as how to run the TinyGo compiler using our Docker container.

## Supported boards/targets

You can compile TinyGo programs for microcontrollers, WebAssembly and Linux.

The following 32 microcontroller boards are currently supported:

* [Adafruit Circuit Playground Bluefruit](https://www.adafruit.com/product/4333)
* [Adafruit Circuit Playground Express](https://www.adafruit.com/product/3333)
* [Adafruit CLUE Alpha](https://www.adafruit.com/product/4500)
* [Adafruit Feather M0](https://www.adafruit.com/product/2772)
* [Adafruit Feather M4](https://www.adafruit.com/product/3857)
* [Adafruit ItsyBitsy M0](https://www.adafruit.com/product/3727)
* [Adafruit ItsyBitsy M4](https://www.adafruit.com/product/3800)
* [Adafruit Metro M4 Express Airlift](https://www.adafruit.com/product/4000)
* [Adafruit PyBadge](https://www.adafruit.com/product/4200)
* [Adafruit PyPortal](https://www.adafruit.com/product/4116)
* [Adafruit Trinket M0](https://www.adafruit.com/product/3500)
* [Arduino Mega 2560](https://store.arduino.cc/arduino-mega-2560-rev3)
* [Arduino Nano](https://store.arduino.cc/arduino-nano)
* [Arduino Nano33 IoT](https://store.arduino.cc/nano-33-iot)
* [Arduino Uno](https://store.arduino.cc/arduino-uno-rev3)
* [BBC micro:bit](https://microbit.org/)
* [Digispark](http://digistump.com/products/1)
* [Game Boy Advance](https://en.wikipedia.org/wiki/Game_Boy_Advance)
* [Makerdiary nRF52840-MDK](https://wiki.makerdiary.com/nrf52840-mdk/)
* [Nordic Semiconductor PCA10031](https://www.nordicsemi.com/eng/Products/nRF51-Dongle)
* [Nordic Semiconductor PCA10040](https://www.nordicsemi.com/eng/Products/Bluetooth-low-energy/nRF52-DK)
* [Nordic Semiconductor PCA10056](https://www.nordicsemi.com/Software-and-Tools/Development-Kits/nRF52840-DK)
* [Particle Argon](https://docs.particle.io/datasheets/wi-fi/argon-datasheet/)
* [Particle Boron](https://docs.particle.io/datasheets/cellular/boron-datasheet/)
* [Particle Xenon](https://docs.particle.io/datasheets/discontinued/xenon-datasheet/)
* [Phytec reel board](https://www.phytec.eu/product-eu/internet-of-things/reelboard/)
* [PineTime DevKit](https://www.pine64.org/pinetime/)
* [SiFIve HiFive1](https://www.sifive.com/boards/hifive1)
* [ST Micro "Nucleo F103RB"](https://www.st.com/en/evaluation-tools/nucleo-f103rb.html)
* [ST Micro STM32F103XX "Bluepill"](http://wiki.stm32duino.com/index.php?title=Blue_Pill)
* [ST Micro STM32F407 "Discovery"](https://www.st.com/en/evaluation-tools/stm32f4discovery.html)
* [X9 Pro smartwatch](https://github.com/curtpw/nRF5x-device-reverse-engineering/tree/master/X9-nrf52832-activity-tracker/)

For more information, see [this list of boards](https://tinygo.org/microcontrollers/). Pull requests for additional support are welcome!

## Currently supported features:

For a description of currently supported Go language features, please see [https://tinygo.org/lang-support/](https://tinygo.org/lang-support/).

## Documentation

Documentation is located on our web site at [https://tinygo.org/](https://tinygo.org/).

You can find the web site code at [https://github.com/tinygo-org/tinygo-site](https://github.com/tinygo-org/tinygo-site).

## Getting help

If you're looking for a more interactive way to discuss TinyGo usage or
development, we have a [#TinyGo channel](https://gophers.slack.com/messages/CDJD3SUP6/)
on the [Gophers Slack](https://gophers.slack.com).

If you need an invitation for the Gophers Slack, you can generate one here which
should arrive fairly quickly (under 1 min): https://invite.slack.golangbridge.org

## Contributing

Your contributions are welcome!

Please take a look at our [CONTRIBUTING.md](./CONTRIBUTING.md) document for details.

## Project Scope

Goals:

* Have very small binary sizes. Don't pay for what you don't use.
* Support for most common microcontroller boards.
* Be usable on the web using WebAssembly.
* Good CGo support, with no more overhead than a regular function call.
* Support most standard library packages and compile most Go code without modification.

Non-goals:

* Using more than one core.
* Be efficient while using zillions of goroutines. However, good goroutine support is certainly a goal.
* Be as fast as `gc`. However, LLVM will probably be better at optimizing certain things so TinyGo might actually turn out to be faster for number crunching.
* Be able to compile every Go program out there.

## Why this project exists

> We never expected Go to be an embedded language and so its got serious problems...

-- Rob Pike, [GopherCon 2014 Opening Keynote](https://www.youtube.com/watch?v=VoS7DsT1rdM&feature=youtu.be&t=2799)

TinyGo is a project to bring Go to microcontrollers and small systems with a single processor core. It is similar to [emgo](https://github.com/ziutek/emgo) but a major difference is that we want to keep the Go memory model (which implies garbage collection of some sort). Another difference is that TinyGo uses LLVM internally instead of emitting C, which hopefully leads to smaller and more efficient code and certainly leads to more flexibility.

The original reasoning was: if [Python](https://micropython.org/) can run on microcontrollers, then certainly [Go](https://golang.org/) should be able to run on even lower level micros.

## License

This project is licensed under the BSD 3-clause license, just like the [Go project](https://golang.org/LICENSE) itself.

Some code has been copied from the LLVM project and is therefore licensed under [a variant of the Apache 2.0 license](http://releases.llvm.org/10.0.0/LICENSE.TXT). This has been clearly indicated in the header of these files.
