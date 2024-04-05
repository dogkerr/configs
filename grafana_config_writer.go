package main

//  ini buat generate config grafana dashboard file
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

type GrafanaConfig struct {
	Dashboard struct {
		Inputs []struct {
			Name        string `json:"name"`
			Label       string `json:"label"`
			Description string `json:"description"`
			Type        string `json:"type"`
			PluginID    string `json:"pluginId"`
			PluginName  string `json:"pluginName"`
		} `json:"__inputs"`
		Requires []struct {
			Type    string `json:"type"`
			ID      string `json:"id"`
			Name    string `json:"name"`
			Version string `json:"version"`
		} `json:"__requires"`
		Annotations struct {
			List []struct {
				BuiltIn    int    `json:"builtIn"`
				Datasource string `json:"datasource"`
				Enable     bool   `json:"enable"`
				Hide       bool   `json:"hide"`
				IconColor  string `json:"iconColor"`
				Name       string `json:"name"`
				Type       string `json:"type"`
			} `json:"list"`
		} `json:"annotations"`
		Description  string        `json:"description"`
		Editable     bool          `json:"editable"`
		GnetID       int           `json:"gnetId"`
		GraphTooltip int           `json:"graphTooltip"`
		ID           interface{}   `json:"id"`
		Iteration    int64         `json:"iteration"`
		Links        []interface{} `json:"links"`
		Panels       []struct {
			CacheTimeout    interface{} `json:"cacheTimeout,omitempty"`
			ColorBackground bool        `json:"colorBackground,omitempty"`
			ColorValue      bool        `json:"colorValue,omitempty"`
			Colors          []string    `json:"colors,omitempty"`
			Datasource      string      `json:"datasource"`
			Decimals        int         `json:"decimals,omitempty"`
			Editable        bool        `json:"editable"`
			Error           bool        `json:"error"`
			Format          string      `json:"format,omitempty"`
			Gauge           struct {
				MaxValue         int  `json:"maxValue"`
				MinValue         int  `json:"minValue"`
				Show             bool `json:"show"`
				ThresholdLabels  bool `json:"thresholdLabels"`
				ThresholdMarkers bool `json:"thresholdMarkers"`
			} `json:"gauge,omitempty"`
			GridPos struct {
				H int `json:"h"`
				W int `json:"w"`
				X int `json:"x"`
				Y int `json:"y"`
			} `json:"gridPos"`
			Height       string        `json:"height,omitempty"`
			ID           int           `json:"id"`
			Interval     interface{}   `json:"interval,omitempty"`
			Links        []interface{} `json:"links"`
			MappingType  int           `json:"mappingType,omitempty"`
			MappingTypes []struct {
				Name  string `json:"name"`
				Value int    `json:"value"`
			} `json:"mappingTypes,omitempty"`
			MaxDataPoints   int         `json:"maxDataPoints,omitempty"`
			NullPointMode   string      `json:"nullPointMode,omitempty"`
			NullText        interface{} `json:"nullText,omitempty"`
			Postfix         string      `json:"postfix,omitempty"`
			PostfixFontSize string      `json:"postfixFontSize,omitempty"`
			Prefix          string      `json:"prefix,omitempty"`
			PrefixFontSize  string      `json:"prefixFontSize,omitempty"`
			RangeMaps       []struct {
				From string `json:"from"`
				Text string `json:"text"`
				To   string `json:"to"`
			} `json:"rangeMaps,omitempty"`
			Sparkline struct {
				FillColor string `json:"fillColor"`
				Full      bool   `json:"full"`
				LineColor string `json:"lineColor"`
				Show      bool   `json:"show"`
			} `json:"sparkline,omitempty"`
			TableColumn string `json:"tableColumn,omitempty"`
			Targets     []struct {
				Expr           string `json:"expr"`
				Format         string `json:"format"`
				Hide           bool   `json:"hide"`
				IntervalFactor int    `json:"intervalFactor"`
				LegendFormat   string `json:"legendFormat"`
				RefID          string `json:"refId"`
				Step           int    `json:"step"`
			} `json:"targets"`
			Thresholds    string `json:"thresholds,omitempty"`
			Title         string `json:"title"`
			Type          string `json:"type"`
			ValueFontSize string `json:"valueFontSize,omitempty"`
			ValueMaps     []struct {
				Op    string `json:"op"`
				Text  string `json:"text"`
				Value string `json:"value"`
			} `json:"valueMaps,omitempty"`
			ValueName   string `json:"valueName,omitempty"`
			AliasColors struct {
				SENT string `json:"SENT"`
			} `json:"aliasColors,omitempty"`
			Bars       bool `json:"bars,omitempty"`
			DashLength int  `json:"dashLength,omitempty"`
			Dashes     bool `json:"dashes,omitempty"`
			Fill       int  `json:"fill,omitempty"`
			Grid       struct {
			} `json:"grid,omitempty"`
			Legend struct {
				Avg     bool `json:"avg"`
				Current bool `json:"current"`
				Max     bool `json:"max"`
				Min     bool `json:"min"`
				Show    bool `json:"show"`
				Total   bool `json:"total"`
				Values  bool `json:"values"`
			} `json:"legend,omitempty"`
			Lines           bool          `json:"lines,omitempty"`
			Linewidth       int           `json:"linewidth,omitempty"`
			Percentage      bool          `json:"percentage,omitempty"`
			Pointradius     int           `json:"pointradius,omitempty"`
			Points          bool          `json:"points,omitempty"`
			Renderer        string        `json:"renderer,omitempty"`
			SeriesOverrides []interface{} `json:"seriesOverrides,omitempty"`
			SpaceLength     int           `json:"spaceLength,omitempty"`
			Stack           bool          `json:"stack,omitempty"`
			SteppedLine     bool          `json:"steppedLine,omitempty"`
			TimeFrom        interface{}   `json:"timeFrom,omitempty"`
			TimeShift       interface{}   `json:"timeShift,omitempty"`
			Tooltip         struct {
				MsResolution bool   `json:"msResolution"`
				Shared       bool   `json:"shared"`
				Sort         int    `json:"sort"`
				ValueType    string `json:"value_type"`
			} `json:"tooltip,omitempty"`
			Transparent bool `json:"transparent,omitempty"`
			Xaxis       struct {
				Buckets interface{}   `json:"buckets"`
				Mode    string        `json:"mode"`
				Name    interface{}   `json:"name"`
				Show    bool          `json:"show"`
				Values  []interface{} `json:"values"`
			} `json:"xaxis,omitempty"`
			Yaxes []struct {
				Format  string      `json:"format"`
				Label   interface{} `json:"label"`
				LogBase int         `json:"logBase"`
				Max     interface{} `json:"max"`
				Min     interface{} `json:"min"`
				Show    bool        `json:"show"`
			} `json:"yaxes,omitempty"`
			Yaxis struct {
				Align      bool        `json:"align"`
				AlignLevel interface{} `json:"alignLevel"`
			} `json:"yaxis,omitempty"`
			Alert struct {
				Conditions []struct {
					Evaluator struct {
						Params []float64 `json:"params"`
						Type   string    `json:"type"`
					} `json:"evaluator"`
					Query struct {
						Params []string `json:"params"`
					} `json:"query"`
					Reducer struct {
						Params []interface{} `json:"params"`
						Type   string        `json:"type"`
					} `json:"reducer"`
					Type string `json:"type"`
				} `json:"conditions"`
				ExecutionErrorState string `json:"executionErrorState"`
				Frequency           string `json:"frequency"`
				Handler             int    `json:"handler"`
				Name                string `json:"name"`
				NoDataState         string `json:"noDataState"`
				Notifications       []struct {
					ID int `json:"id"`
				} `json:"notifications"`
			} `json:"alert,omitempty"`
			Columns []struct {
				Text  string `json:"text"`
				Value string `json:"value"`
			} `json:"columns,omitempty"`
			FontSize   string      `json:"fontSize,omitempty"`
			PageSize   interface{} `json:"pageSize,omitempty"`
			Scroll     bool        `json:"scroll,omitempty"`
			ShowHeader bool        `json:"showHeader,omitempty"`
			Sort       struct {
				Col  int  `json:"col"`
				Desc bool `json:"desc"`
			} `json:"sort,omitempty"`
			Styles []struct {
				ColorMode  interface{} `json:"colorMode"`
				Colors     []string    `json:"colors"`
				Decimals   int         `json:"decimals"`
				Pattern    string      `json:"pattern"`
				Thresholds []string    `json:"thresholds"`
				Type       string      `json:"type"`
				Unit       string      `json:"unit"`
			} `json:"styles,omitempty"`
			Transform string `json:"transform,omitempty"`
		} `json:"panels"`
		Refresh       string        `json:"refresh"`
		SchemaVersion int           `json:"schemaVersion"`
		Style         string        `json:"style"`
		Tags          []interface{} `json:"tags"`
		Templating    struct {
			List []struct {
				AllValue string `json:"allValue,omitempty"`
				Current  struct {
				} `json:"current"`
				Datasource     string        `json:"datasource"`
				Hide           int           `json:"hide"`
				IncludeAll     bool          `json:"includeAll"`
				Label          string        `json:"label"`
				Multi          bool          `json:"multi"`
				Name           string        `json:"name"`
				Options        []interface{} `json:"options"`
				Query          string        `json:"query"`
				Refresh        int           `json:"refresh"`
				Regex          string        `json:"regex,omitempty"`
				Sort           int           `json:"sort,omitempty"`
				TagValuesQuery interface{}   `json:"tagValuesQuery,omitempty"`
				Tags           []interface{} `json:"tags,omitempty"`
				TagsQuery      interface{}   `json:"tagsQuery,omitempty"`
				Type           string        `json:"type"`
				UseTags        bool          `json:"useTags,omitempty"`
				Auto           bool          `json:"auto,omitempty"`
				AutoCount      int           `json:"auto_count,omitempty"`
				AutoMin        string        `json:"auto_min,omitempty"`
			} `json:"list"`
		} `json:"templating"`
		Time struct {
			From string `json:"from"`
			To   string `json:"to"`
		} `json:"time"`
		Timepicker struct {
			RefreshIntervals []string `json:"refresh_intervals"`
			TimeOptions      []string `json:"time_options"`
		} `json:"timepicker"`
		Timezone string `json:"timezone"`
		Title    string `json:"title"`
		UID      string `json:"uid"`
		Version  int    `json:"version"`
	} `json:"dashboard"`
}
type DataSourceResponse struct {
	ID    string `json:"id"`
	Uid   string `json:"uid"`
	OrgId string `json:"orgId"`
	Type  string `json:"type"`
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[random.Intn(len(charset))]
	}
	return string(result)
}

func createNewDashboardPerUser(userId string) string {
	var client = &http.Client{}
	var data []DataSourceResponse

	request, err := http.NewRequest("GET", "http://localhost:3000/api/datasources", nil)
	if err != nil {
		fmt.Errorf("Error: " + err.Error())
	}

	response, err := client.Do(request)
	if err != nil {
		fmt.Errorf("Error: " + err.Error())
	}

	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		fmt.Errorf("Error: " + err.Error())
	}

	// Open our jsonFile
	jsonFile, err := os.Open("docker-quest-prometheus.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var grafanaConfig GrafanaConfig

	json.Unmarshal(byteValue, &grafanaConfig)

	fmt.Println("hasil data source get: ")
	// fmt.Println(datasourceId) // salah disini
	datasourceId := ""
	for _, ds := range data {
		fmt.Println(ds)
		fmt.Println(ds.Type)
		if ds.Type == "prometheus" {
			fmt.Println("prome: ", ds)
			datasourceId = ds.Uid
		}
	}
	fmt.Println(datasourceId)

	grafanaConfig.Dashboard.Inputs[0].PluginID = datasourceId
	for i, _ := range grafanaConfig.Dashboard.Panels {
		grafanaConfig.Dashboard.Panels[i].Datasource = datasourceId
	}

	for i, _ := range grafanaConfig.Dashboard.Templating.List {
		grafanaConfig.Dashboard.Templating.List[i].Datasource = datasourceId
	}

	// idxCpuUsage := 14
	// idxSentNetworkTraffic := 13
	// idxRcvdNetworkTraffic := 12
	// idxSwap := 15
	// idxMemUsage :=16
	// idxLimitMemory := 17
	// idxUsageMemory := 18
	// idxRemainingMemory=19
	user := userId
	cpuUsageUser := grafanaConfig.Dashboard.Panels[14].Targets[0].Expr
	fmt.Println(cpuUsageUser)
	cpuUsageByUser := "sum(rate(container_cpu_usage_seconds_total{container_label_user_id=~\"" + user + "\"}[$interval])) by (name) * 100"
	grafanaConfig.Dashboard.Panels[14].Targets[0].Expr = cpuUsageByUser

	sentNetworkTrafficUser := grafanaConfig.Dashboard.Panels[13].Targets[0].Expr
	fmt.Println(sentNetworkTrafficUser)
	sentNetworkTrafficByUser := "sum(rate(container_network_transmit_bytes_total{container_label_user_id=~\"" + user + " \"}[$interval])) by (name)"
	fmt.Println(sentNetworkTrafficByUser)
	grafanaConfig.Dashboard.Panels[13].Targets[0].Expr = sentNetworkTrafficByUser

	rcvdNetworkTrafficUser := grafanaConfig.Dashboard.Panels[12].Targets[0].Expr
	fmt.Println(rcvdNetworkTrafficUser)
	// sum(rate(container_network_receive_bytes_total{container_label_user_id=~\"18d2e020-538d-449a-8e9c-02e4e5cf41111\"}[$interval])) by (name)
	grafanaConfig.Dashboard.Panels[12].Targets[0].Expr = "sum(rate(container_network_receive_bytes_total{container_label_user_id=~\"" + user + "\"}[$interval])) by (name)"

	swapUser := grafanaConfig.Dashboard.Panels[15].Targets[0].Expr
	fmt.Println(swapUser)
	// sum(container_memory_swap{name=~\".+\"}) by (name)
	grafanaConfig.Dashboard.Panels[15].Targets[0].Expr = "sum(container_memory_swap{container_label_user_id=~\"" + user + "\"}) by (name)"

	memUsage := grafanaConfig.Dashboard.Panels[16].Targets[0].Expr
	fmt.Println(memUsage)
	grafanaConfig.Dashboard.Panels[16].Targets[0].Expr = "sum(container_memory_rss{container_label_user_id=~\"" + user + "\"}) by (name)"
	grafanaConfig.Dashboard.Panels[16].Targets[1].Expr = "container_memory_usage_bytes{container_label_user_id=~\"" + user + "\"}"

	limitMemory := grafanaConfig.Dashboard.Panels[17].Targets[0].Expr
	fmt.Println(limitMemory)
	grafanaConfig.Dashboard.Panels[17].Targets[0].Expr = "sum(container_spec_memory_limit_bytes{container_label_user_id=~\"" + user + "\"} - container_memory_usage_bytes{container_label_user_id=~\"" + user + "\"}) by (name) "

	usageMemory := grafanaConfig.Dashboard.Panels[18].Targets[0].Expr
	fmt.Println(usageMemory)
	grafanaConfig.Dashboard.Panels[18].Targets[2].Expr = " container_memory_usage_bytes{container_label_user_id=~\"" + user + "\"} "

	grafanaConfig.Dashboard.Panels[1].Targets[0].Expr = "count(rate(container_last_seen{container_label_user_id=~\"" + user + "\"}[$interval]))"

	// remainingMemory := grafanaConfig.Dashboard.Panels[19].Targets[0].Expr
	// fmt.Println(remainingMemory)
	// grafanaConfig.Dashboard.Panels[19].Targets[0].Expr = "sum(100 - ((container_spec_memory_limit_bytes{container_label_user_id=~\"" + user + "\"} - container_memory_usage_bytes{container_label_user_id=~\"" + user + "\"})  * 100 / container_spec_memory_limit_bytes{container_label_user_id=~\"" + user + "\"}) ) by (name)"

	grafanaConfig.Dashboard.Panels = grafanaConfig.Dashboard.Panels[:len(grafanaConfig.Dashboard.Panels)-1]

	randomString := generateRandomString(8)
	grafanaConfig.Dashboard.Title = randomString
	grafanaConfig.Dashboard.UID = randomString
	grafanaConfig.Dashboard.Style = "light"

	file, _ := json.MarshalIndent(grafanaConfig, "", " ")
	// _ = ioutil.WriteFile(randomString+".json", file, 0644)

	createDashboardUrl := "http://localhost:3000/api/dashboards/db"
	r, err := http.NewRequest("POST", createDashboardUrl, bytes.NewBuffer(file))
	if err != nil {
		fmt.Println(err)
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", "Bearer glsa_i11OKMoEa9AsG5mCMNBRb4PWU7dmRiX5_94402364")
	client = &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	fmt.Println("response create dashboard: " + res.Status)
	return randomString
}

func main() {
	uidDashboard := createNewDashboardPerUser("18d2e020-538d-449a-8e9c-02e4e5cf41111")

	fmt.Println("http://127.0.0.1/d-solo/" + uidDashboard + "/" + strings.ToLower(uidDashboard) + "?orgId=1&refresh=5s&from=" + "now-5m" + "&theme=light&to=" + "now" + "&panelId=8")
	fmt.Println("http://127.0.0.1/d-solo/" + uidDashboard + "/" + strings.ToLower(uidDashboard) + "?orgId=1&refresh=5s&from=" + "now-5m" + "&theme=light&to=" + "now" + "&panelId=9")
	fmt.Println("http://127.0.0.1/d-solo/" + uidDashboard + "/" + strings.ToLower(uidDashboard) + "?orgId=1&refresh=5s&from=" + "now-5m" + "&theme=light&to=" + "now" + "&panelId=1")
	fmt.Println("http://127.0.0.1/d-solo/" + uidDashboard + "/" + strings.ToLower(uidDashboard) + "?orgId=1&refresh=5s&from=" + "now-5m" + "&theme=light&to=" + "now" + "&panelId=34")
	fmt.Println("http://127.0.0.1/d-solo/" + uidDashboard + "/" + strings.ToLower(uidDashboard) + "?orgId=1&refresh=5s&from=" + "now-5m" + "&theme=light&to=" + "now" + "&panelId=10")

	fmt.Println("http://127.0.0.1/d-solo/" + uidDashboard + "/" + strings.ToLower(uidDashboard) + "?orgId=1&refresh=5s&from=" + "now-5m" + "&theme=light&to=" + "now" + "&panelId=37")

	fmt.Println("http://127.0.0.1/d-solo/" + uidDashboard + "/" + strings.ToLower(uidDashboard) + "?orgId=1&refresh=5s&from=" + "now-5m" + "&theme=light&to=" + "now" + "&panelId=5")

	fmt.Println("http://127.0.0.1/d-solo/" + uidDashboard + "/" + strings.ToLower(uidDashboard) + "?orgId=1&refresh=5s&from=" + "now-5m" + "&theme=light&to=" + "now" + "&panelId=31")

	// sentNetworkTraffByUser =
}
