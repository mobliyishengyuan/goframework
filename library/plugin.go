package library

const (
	PluginPreOptParse = iota
	PluginPostOptParse
	PluginPreServerStart
)

const (
	LenPluginList = 10
)

type plugin func()

var pluginMap map[int]([]plugin)

func init() {
	pluginMap = make(map[int]([]plugin))
	
	pluginMap[PluginPreOptParse] = make([]plugin, LenPluginList)
	pluginMap[PluginPostOptParse] = make([]plugin, LenPluginList)
	pluginMap[PluginPreServerStart] = make([]plugin, LenPluginList)
}

func RegisterPlugin(time int, callback plugin) {
	pluginMap[time] = append(pluginMap[time], callback)
}

func QueryPlugin(time int) {
	pluginList, ok := pluginMap[time]
	
	if !ok {
		return
	}
	
	for _, v := range pluginList {
		
		if v == nil {
			continue;
		}
		
		v()
	}
}