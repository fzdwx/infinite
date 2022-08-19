# input

:::tip
input 是一个用于接收用户输入的基础组件。
:::

## original demo

这是一个最原始的demo，直接使用`startUp` 启动 `input`。

![demo](https://user-images.githubusercontent.com/65269574/185384071-1d383fe9-fc28-4abc-904f-8a2de524c29f.gif)

:::details 代码
<<< @/_examples/input-original-demo/main.go
:::

## input text

接下来就是 `input text` 这个组件了，之所以直接介绍这一个，是因为它完全就是 `input`
的功能的封装，只是把可以修改的配置用 [`option`](https://github.com/fzdwx/infinite/blob/main/components/input/text/options.go#L9)
的形式暴露了出来，并且修改了一些 [默认配置](https://github.com/fzdwx/infinite/blob/main/components/input/text/text.go#L16)
，稍微改变了一下样式， 当然也提供了没有加样式的版本，你只需要添加一个`text.WithPure()`这个option。

![demo](https://user-images.githubusercontent.com/65269574/185403896-a9ab46c0-9f3d-4957-9c30-a336c9ba855c.gif)

::: details 代码
<<< @/_examples/input-text/main.go
:::

## options

> 这里是可供修改的配置项，你可以试着将它变成你的形状。

### 修改显示区域

:::tip
`Focus` 就代表当前正在运行这个组件会显示，`UnFocus` 表示组件退出使用时才会显示。
:::

`input text` 主要显示的区域有以下几个部分:

- `FocusSymbol/UnFocusSymbol`: 提示符号
- `Prompt`: 提示信息
- `FocusInterval/UnFocusInterval`: 间隔符号
- `DefaultValue/Text`: 默认值与实际值
- `Cursor`: 光标

它们在 `input text` 中的大概位置:
![image](https://user-images.githubusercontent.com/65269574/185407799-c332bc4a-91b7-48d8-85fa-b55bd20454e2.png)

你可以向这样修改它们，并且可以修改它们的`style`。

<<< @/_examples/input-text/optiondemo.txt{2-6 go}

### keymap

- `Confirm`: 用户确认输入, 默认是 `enter`。
- `Quit`:  用户强制退出程序, 默认是 `ctrl+c`。

<<< @/_examples/input-text/option_keymap.txt{3-9 go}

### required

你可以通过设置 `required` , 让用户必须输入一些东西。

- `Required`: 开启强制输入
- `RequiredMsg`: 用户 `Confirm` 了却没有输入任何东西就会打印出来。
- `RequiredMsgKeepAliveTime`: `RequiredMsg` 的存活时间。

<<< @/_examples/input-text/option_required.txt{2-4 go}

### default value

设置 `DefaultValue`, 即用户直接按下`Confirm`就会返回该值。

:::tip
如果设置了 `Required` ,`DefaultValue` 会被忽略。
:::

<<< @/_examples/input-text/defaultval.txt{2 go}

### disableOutputResult

该选项是用于禁用输出结果。

<<< @/_examples/input-text/disableOutputResult.txt{2 go}

禁用前:
![demo](https://user-images.githubusercontent.com/65269574/185613513-567b933c-ee9b-42fb-9e09-e6ec37ee46c5.gif)

禁用后:
![demo](https://user-images.githubusercontent.com/65269574/185613335-d186edfc-65e0-4045-8ccb-772c559791ee.gif)


### 改变text的显示模式

`text` 默认是输入什么就显示什么，你可以通过下面的option来修改它。

#### 显示空白

![demo](https://user-images.githubusercontent.com/65269574/185615086-5b6d2620-7594-4de4-8f43-d31f8461ddc9.gif)

::: details 代码
<<< @/_examples/input-text/echonone.txt{12 go}
:::

#### 显示自定义字符

![demo](https://user-images.githubusercontent.com/65269574/185615753-fd095426-cf19-40ed-b29f-229400e328dd.gif)

::: details 代码
<<< @/_examples/input-text/echopwd.txt{13 go}
:::