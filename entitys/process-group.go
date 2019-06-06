package entitys

type Permissions struct {
	CanRead  bool `json:"canRead"`
	CanWrite bool `json:"canWrite"`
}

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Revision struct {
	ClientID string `json:"clientId"`
	Version  int    `json:"version"`
}

type ProcessGroup struct {
	Revision Revision `json:"revision"`
	ID       string   `json:"id"`
	URI      string   `json:"uri"`
	Position struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	} `json:"position"`
	Permissions struct {
		CanRead  bool `json:"canRead"`
		CanWrite bool `json:"canWrite"`
	} `json:"permissions"`
	Bulletins []interface{} `json:"bulletins"`
	Component struct {
		ID            string `json:"id"`
		ParentGroupID string `json:"parentGroupId"`
		Position      struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"position"`
		Name      string `json:"name"`
		Comments  string `json:"comments"`
		Variables struct {
		} `json:"variables"`
		RunningCount                 int `json:"runningCount"`
		StoppedCount                 int `json:"stoppedCount"`
		InvalidCount                 int `json:"invalidCount"`
		DisabledCount                int `json:"disabledCount"`
		ActiveRemotePortCount        int `json:"activeRemotePortCount"`
		InactiveRemotePortCount      int `json:"inactiveRemotePortCount"`
		UpToDateCount                int `json:"upToDateCount"`
		LocallyModifiedCount         int `json:"locallyModifiedCount"`
		StaleCount                   int `json:"staleCount"`
		LocallyModifiedAndStaleCount int `json:"locallyModifiedAndStaleCount"`
		SyncFailureCount             int `json:"syncFailureCount"`
		InputPortCount               int `json:"inputPortCount"`
		OutputPortCount              int `json:"outputPortCount"`
	} `json:"component"`
	Status struct {
		ID                 string `json:"id"`
		Name               string `json:"name"`
		StatsLastRefreshed string `json:"statsLastRefreshed"`
		AggregateSnapshot  struct {
			ID                    string `json:"id"`
			Name                  string `json:"name"`
			FlowFilesIn           int    `json:"flowFilesIn"`
			BytesIn               int    `json:"bytesIn"`
			Input                 string `json:"input"`
			FlowFilesQueued       int    `json:"flowFilesQueued"`
			BytesQueued           int    `json:"bytesQueued"`
			Queued                string `json:"queued"`
			QueuedCount           string `json:"queuedCount"`
			QueuedSize            string `json:"queuedSize"`
			BytesRead             int    `json:"bytesRead"`
			Read                  string `json:"read"`
			BytesWritten          int    `json:"bytesWritten"`
			Written               string `json:"written"`
			FlowFilesOut          int    `json:"flowFilesOut"`
			BytesOut              int    `json:"bytesOut"`
			Output                string `json:"output"`
			FlowFilesTransferred  int    `json:"flowFilesTransferred"`
			BytesTransferred      int    `json:"bytesTransferred"`
			Transferred           string `json:"transferred"`
			BytesReceived         int    `json:"bytesReceived"`
			FlowFilesReceived     int    `json:"flowFilesReceived"`
			Received              string `json:"received"`
			BytesSent             int    `json:"bytesSent"`
			FlowFilesSent         int    `json:"flowFilesSent"`
			Sent                  string `json:"sent"`
			ActiveThreadCount     int    `json:"activeThreadCount"`
			TerminatedThreadCount int    `json:"terminatedThreadCount"`
		} `json:"aggregateSnapshot"`
	} `json:"status"`
	RunningCount                 int `json:"runningCount"`
	StoppedCount                 int `json:"stoppedCount"`
	InvalidCount                 int `json:"invalidCount"`
	DisabledCount                int `json:"disabledCount"`
	ActiveRemotePortCount        int `json:"activeRemotePortCount"`
	InactiveRemotePortCount      int `json:"inactiveRemotePortCount"`
	UpToDateCount                int `json:"upToDateCount"`
	LocallyModifiedCount         int `json:"locallyModifiedCount"`
	StaleCount                   int `json:"staleCount"`
	LocallyModifiedAndStaleCount int `json:"locallyModifiedAndStaleCount"`
	SyncFailureCount             int `json:"syncFailureCount"`
	InputPortCount               int `json:"inputPortCount"`
	OutputPortCount              int `json:"outputPortCount"`
}

type ProcessorComponent struct {
	ID            string   `json:"id"`
	ParentGroupID string   `json:"parentGroupId"`
	Position      Position `json:"position"`
	Name          string   `json:"name"`
	Type          string   `json:"type"`
	Bundle        Bundle   `json:"bundle"`
	State         string   `json:"state"`
	Style         struct {
	} `json:"style"`
	Relationships []struct {
		Name          string `json:"name"`
		Description   string `json:"description"`
		AutoTerminate bool   `json:"autoTerminate"`
	} `json:"relationships"`
	SupportsParallelProcessing bool   `json:"supportsParallelProcessing"`
	SupportsEventDriven        bool   `json:"supportsEventDriven"`
	SupportsBatching           bool   `json:"supportsBatching"`
	PersistsState              bool   `json:"persistsState"`
	Restricted                 bool   `json:"restricted"`
	Deprecated                 bool   `json:"deprecated"`
	ExecutionNodeRestricted    bool   `json:"executionNodeRestricted"`
	MultipleVersionsAvailable  bool   `json:"multipleVersionsAvailable"`
	InputRequirement           string `json:"inputRequirement"`
	Config                     struct {
		Properties struct {
			DatabaseConnectionPoolingService string      `json:"Database Connection Pooling Service"`
			DbFetchDbType                    string      `json:"db-fetch-db-type"`
			TableName                        string      `json:"Table Name"`
			ColumnsToReturn                  interface{} `json:"Columns to Return"`
			DbFetchWhereClause               interface{} `json:"db-fetch-where-clause"`
			DbFetchSQLQuery                  string      `json:"db-fetch-sql-query"`
			QdbtrRecordWriter                string      `json:"qdbtr-record-writer"`
			MaximumValueColumns              string      `json:"Maximum-value Columns"`
			MaxWaitTime                      string      `json:"Max Wait Time"`
			FetchSize                        string      `json:"Fetch Size"`
			QdbtMaxRows                      string      `json:"qdbt-max-rows"`
			QdbtOutputBatchSize              string      `json:"qdbt-output-batch-size"`
			QdbtMaxFrags                     string      `json:"qdbt-max-frags"`
			QdbtrNormalize                   string      `json:"qdbtr-normalize"`
			DbfUserLogicalTypes              string      `json:"dbf-user-logical-types"`
		} `json:"properties"`
		Descriptors struct {
			DatabaseConnectionPoolingService struct {
				Name            string `json:"name"`
				DisplayName     string `json:"displayName"`
				Description     string `json:"description"`
				AllowableValues []struct {
					AllowableValue struct {
						DisplayName string `json:"displayName"`
						Value       string `json:"value"`
					} `json:"allowableValue"`
					CanRead bool `json:"canRead"`
				} `json:"allowableValues"`
				Required                          bool   `json:"required"`
				Sensitive                         bool   `json:"sensitive"`
				Dynamic                           bool   `json:"dynamic"`
				SupportsEl                        bool   `json:"supportsEl"`
				ExpressionLanguageScope           string `json:"expressionLanguageScope"`
				IdentifiesControllerService       string `json:"identifiesControllerService"`
				IdentifiesControllerServiceBundle struct {
					Group    string `json:"group"`
					Artifact string `json:"artifact"`
					Version  string `json:"version"`
				} `json:"identifiesControllerServiceBundle"`
			} `json:"Database Connection Pooling Service"`
			DbFetchDbType struct {
				Name            string `json:"name"`
				DisplayName     string `json:"displayName"`
				Description     string `json:"description"`
				DefaultValue    string `json:"defaultValue"`
				AllowableValues []struct {
					AllowableValue struct {
						DisplayName string `json:"displayName"`
						Value       string `json:"value"`
						Description string `json:"description"`
					} `json:"allowableValue"`
					CanRead bool `json:"canRead"`
				} `json:"allowableValues"`
				Required                bool   `json:"required"`
				Sensitive               bool   `json:"sensitive"`
				Dynamic                 bool   `json:"dynamic"`
				SupportsEl              bool   `json:"supportsEl"`
				ExpressionLanguageScope string `json:"expressionLanguageScope"`
			} `json:"db-fetch-db-type"`
			TableName struct {
				Name                    string `json:"name"`
				DisplayName             string `json:"displayName"`
				Description             string `json:"description"`
				Required                bool   `json:"required"`
				Sensitive               bool   `json:"sensitive"`
				Dynamic                 bool   `json:"dynamic"`
				SupportsEl              bool   `json:"supportsEl"`
				ExpressionLanguageScope string `json:"expressionLanguageScope"`
			} `json:"Table Name"`
			ColumnsToReturn struct {
				Name                    string `json:"name"`
				DisplayName             string `json:"displayName"`
				Description             string `json:"description"`
				Required                bool   `json:"required"`
				Sensitive               bool   `json:"sensitive"`
				Dynamic                 bool   `json:"dynamic"`
				SupportsEl              bool   `json:"supportsEl"`
				ExpressionLanguageScope string `json:"expressionLanguageScope"`
			} `json:"Columns to Return"`
			DbFetchWhereClause struct {
				Name                    string `json:"name"`
				DisplayName             string `json:"displayName"`
				Description             string `json:"description"`
				Required                bool   `json:"required"`
				Sensitive               bool   `json:"sensitive"`
				Dynamic                 bool   `json:"dynamic"`
				SupportsEl              bool   `json:"supportsEl"`
				ExpressionLanguageScope string `json:"expressionLanguageScope"`
			} `json:"db-fetch-where-clause"`
			DbFetchSQLQuery struct {
				Name                    string `json:"name"`
				DisplayName             string `json:"displayName"`
				Description             string `json:"description"`
				Required                bool   `json:"required"`
				Sensitive               bool   `json:"sensitive"`
				Dynamic                 bool   `json:"dynamic"`
				SupportsEl              bool   `json:"supportsEl"`
				ExpressionLanguageScope string `json:"expressionLanguageScope"`
			} `json:"db-fetch-sql-query"`
			QdbtrRecordWriter struct {
				Name            string `json:"name"`
				DisplayName     string `json:"displayName"`
				Description     string `json:"description"`
				AllowableValues []struct {
					AllowableValue struct {
						DisplayName string `json:"displayName"`
						Value       string `json:"value"`
					} `json:"allowableValue"`
					CanRead bool `json:"canRead"`
				} `json:"allowableValues"`
				Required                          bool   `json:"required"`
				Sensitive                         bool   `json:"sensitive"`
				Dynamic                           bool   `json:"dynamic"`
				SupportsEl                        bool   `json:"supportsEl"`
				ExpressionLanguageScope           string `json:"expressionLanguageScope"`
				IdentifiesControllerService       string `json:"identifiesControllerService"`
				IdentifiesControllerServiceBundle struct {
					Group    string `json:"group"`
					Artifact string `json:"artifact"`
					Version  string `json:"version"`
				} `json:"identifiesControllerServiceBundle"`
			} `json:"qdbtr-record-writer"`
			MaximumValueColumns struct {
				Name                    string `json:"name"`
				DisplayName             string `json:"displayName"`
				Description             string `json:"description"`
				Required                bool   `json:"required"`
				Sensitive               bool   `json:"sensitive"`
				Dynamic                 bool   `json:"dynamic"`
				SupportsEl              bool   `json:"supportsEl"`
				ExpressionLanguageScope string `json:"expressionLanguageScope"`
			} `json:"Maximum-value Columns"`
			MaxWaitTime struct {
				Name                    string `json:"name"`
				DisplayName             string `json:"displayName"`
				Description             string `json:"description"`
				DefaultValue            string `json:"defaultValue"`
				Required                bool   `json:"required"`
				Sensitive               bool   `json:"sensitive"`
				Dynamic                 bool   `json:"dynamic"`
				SupportsEl              bool   `json:"supportsEl"`
				ExpressionLanguageScope string `json:"expressionLanguageScope"`
			} `json:"Max Wait Time"`
			FetchSize struct {
				Name                    string `json:"name"`
				DisplayName             string `json:"displayName"`
				Description             string `json:"description"`
				DefaultValue            string `json:"defaultValue"`
				Required                bool   `json:"required"`
				Sensitive               bool   `json:"sensitive"`
				Dynamic                 bool   `json:"dynamic"`
				SupportsEl              bool   `json:"supportsEl"`
				ExpressionLanguageScope string `json:"expressionLanguageScope"`
			} `json:"Fetch Size"`
			QdbtMaxRows struct {
				Name                    string `json:"name"`
				DisplayName             string `json:"displayName"`
				Description             string `json:"description"`
				DefaultValue            string `json:"defaultValue"`
				Required                bool   `json:"required"`
				Sensitive               bool   `json:"sensitive"`
				Dynamic                 bool   `json:"dynamic"`
				SupportsEl              bool   `json:"supportsEl"`
				ExpressionLanguageScope string `json:"expressionLanguageScope"`
			} `json:"qdbt-max-rows"`
			QdbtOutputBatchSize struct {
				Name                    string `json:"name"`
				DisplayName             string `json:"displayName"`
				Description             string `json:"description"`
				DefaultValue            string `json:"defaultValue"`
				Required                bool   `json:"required"`
				Sensitive               bool   `json:"sensitive"`
				Dynamic                 bool   `json:"dynamic"`
				SupportsEl              bool   `json:"supportsEl"`
				ExpressionLanguageScope string `json:"expressionLanguageScope"`
			} `json:"qdbt-output-batch-size"`
			QdbtMaxFrags struct {
				Name                    string `json:"name"`
				DisplayName             string `json:"displayName"`
				Description             string `json:"description"`
				DefaultValue            string `json:"defaultValue"`
				Required                bool   `json:"required"`
				Sensitive               bool   `json:"sensitive"`
				Dynamic                 bool   `json:"dynamic"`
				SupportsEl              bool   `json:"supportsEl"`
				ExpressionLanguageScope string `json:"expressionLanguageScope"`
			} `json:"qdbt-max-frags"`
			QdbtrNormalize struct {
				Name            string `json:"name"`
				DisplayName     string `json:"displayName"`
				Description     string `json:"description"`
				DefaultValue    string `json:"defaultValue"`
				AllowableValues []struct {
					AllowableValue struct {
						DisplayName string `json:"displayName"`
						Value       string `json:"value"`
					} `json:"allowableValue"`
					CanRead bool `json:"canRead"`
				} `json:"allowableValues"`
				Required                bool   `json:"required"`
				Sensitive               bool   `json:"sensitive"`
				Dynamic                 bool   `json:"dynamic"`
				SupportsEl              bool   `json:"supportsEl"`
				ExpressionLanguageScope string `json:"expressionLanguageScope"`
			} `json:"qdbtr-normalize"`
			DbfUserLogicalTypes struct {
				Name            string `json:"name"`
				DisplayName     string `json:"displayName"`
				Description     string `json:"description"`
				DefaultValue    string `json:"defaultValue"`
				AllowableValues []struct {
					AllowableValue struct {
						DisplayName string `json:"displayName"`
						Value       string `json:"value"`
					} `json:"allowableValue"`
					CanRead bool `json:"canRead"`
				} `json:"allowableValues"`
				Required                bool   `json:"required"`
				Sensitive               bool   `json:"sensitive"`
				Dynamic                 bool   `json:"dynamic"`
				SupportsEl              bool   `json:"supportsEl"`
				ExpressionLanguageScope string `json:"expressionLanguageScope"`
			} `json:"dbf-user-logical-types"`
		} `json:"descriptors"`
		SchedulingPeriod                 string `json:"schedulingPeriod"`
		SchedulingStrategy               string `json:"schedulingStrategy"`
		ExecutionNode                    string `json:"executionNode"`
		PenaltyDuration                  string `json:"penaltyDuration"`
		YieldDuration                    string `json:"yieldDuration"`
		BulletinLevel                    string `json:"bulletinLevel"`
		RunDurationMillis                int    `json:"runDurationMillis"`
		ConcurrentlySchedulableTaskCount int    `json:"concurrentlySchedulableTaskCount"`
		Comments                         string `json:"comments"`
		LossTolerant                     bool   `json:"lossTolerant"`
		DefaultConcurrentTasks           struct {
			TIMERDRIVEN string `json:"TIMER_DRIVEN"`
			EVENTDRIVEN string `json:"EVENT_DRIVEN"`
			CRONDRIVEN  string `json:"CRON_DRIVEN"`
		} `json:"defaultConcurrentTasks"`
		DefaultSchedulingPeriod struct {
			TIMERDRIVEN string `json:"TIMER_DRIVEN"`
			CRONDRIVEN  string `json:"CRON_DRIVEN"`
		} `json:"defaultSchedulingPeriod"`
	} `json:"config"`
	ValidationStatus string `json:"validationStatus"`
	ExtensionMissing bool   `json:"extensionMissing"`
}

type Processor struct {
	Revision struct {
		ClientID string `json:"clientId"`
		Version  int    `json:"version"`
	} `json:"revision"`
	ID          string   `json:"id"`
	URI         string   `json:"uri"`
	Position    Position `json:"position"`
	Permissions struct {
		CanRead  bool `json:"canRead"`
		CanWrite bool `json:"canWrite"`
	} `json:"permissions"`
	Bulletins        []interface{}      `json:"bulletins"`
	Component        ProcessorComponent `json:"component"`
	InputRequirement string             `json:"inputRequirement"`
	Status           struct {
		GroupID            string `json:"groupId"`
		ID                 string `json:"id"`
		Name               string `json:"name"`
		RunStatus          string `json:"runStatus"`
		StatsLastRefreshed string `json:"statsLastRefreshed"`
		AggregateSnapshot  struct {
			ID                    string `json:"id"`
			GroupID               string `json:"groupId"`
			Name                  string `json:"name"`
			Type                  string `json:"type"`
			RunStatus             string `json:"runStatus"`
			ExecutionNode         string `json:"executionNode"`
			BytesRead             int    `json:"bytesRead"`
			BytesWritten          int    `json:"bytesWritten"`
			Read                  string `json:"read"`
			Written               string `json:"written"`
			FlowFilesIn           int    `json:"flowFilesIn"`
			BytesIn               int    `json:"bytesIn"`
			Input                 string `json:"input"`
			FlowFilesOut          int    `json:"flowFilesOut"`
			BytesOut              int    `json:"bytesOut"`
			Output                string `json:"output"`
			TaskCount             int    `json:"taskCount"`
			TasksDurationNanos    int    `json:"tasksDurationNanos"`
			Tasks                 string `json:"tasks"`
			TasksDuration         string `json:"tasksDuration"`
			ActiveThreadCount     int    `json:"activeThreadCount"`
			TerminatedThreadCount int    `json:"terminatedThreadCount"`
		} `json:"aggregateSnapshot"`
	} `json:"status"`
	OperatePermissions struct {
		CanRead  bool `json:"canRead"`
		CanWrite bool `json:"canWrite"`
	} `json:"operatePermissions"`
}

type Connection struct {
	Revision struct {
		Version int `json:"version"`
	} `json:"revision"`
	ID          string `json:"id"`
	URI         string `json:"uri"`
	Permissions struct {
		CanRead  bool `json:"canRead"`
		CanWrite bool `json:"canWrite"`
	} `json:"permissions"`
	Component struct {
		ID            string `json:"id"`
		ParentGroupID string `json:"parentGroupId"`
		Source        struct {
			ID       string `json:"id"`
			Type     string `json:"type"`
			GroupID  string `json:"groupId"`
			Name     string `json:"name"`
			Running  bool   `json:"running"`
			Comments string `json:"comments"`
		} `json:"source"`
		Destination struct {
			ID       string `json:"id"`
			Type     string `json:"type"`
			GroupID  string `json:"groupId"`
			Name     string `json:"name"`
			Running  bool   `json:"running"`
			Comments string `json:"comments"`
		} `json:"destination"`
		Name                          string        `json:"name"`
		LabelIndex                    int           `json:"labelIndex"`
		ZIndex                        int           `json:"zIndex"`
		SelectedRelationships         []string      `json:"selectedRelationships"`
		AvailableRelationships        []string      `json:"availableRelationships"`
		BackPressureObjectThreshold   int           `json:"backPressureObjectThreshold"`
		BackPressureDataSizeThreshold string        `json:"backPressureDataSizeThreshold"`
		FlowFileExpiration            string        `json:"flowFileExpiration"`
		Prioritizers                  []interface{} `json:"prioritizers"`
		Bends                         []interface{} `json:"bends"`
		LoadBalanceStrategy           string        `json:"loadBalanceStrategy"`
		LoadBalancePartitionAttribute string        `json:"loadBalancePartitionAttribute"`
		LoadBalanceCompression        string        `json:"loadBalanceCompression"`
		LoadBalanceStatus             string        `json:"loadBalanceStatus"`
	} `json:"component"`
	Status struct {
		ID                 string `json:"id"`
		GroupID            string `json:"groupId"`
		Name               string `json:"name"`
		StatsLastRefreshed string `json:"statsLastRefreshed"`
		SourceID           string `json:"sourceId"`
		SourceName         string `json:"sourceName"`
		DestinationID      string `json:"destinationId"`
		DestinationName    string `json:"destinationName"`
		AggregateSnapshot  struct {
			ID              string `json:"id"`
			GroupID         string `json:"groupId"`
			Name            string `json:"name"`
			SourceName      string `json:"sourceName"`
			DestinationName string `json:"destinationName"`
			FlowFilesIn     int    `json:"flowFilesIn"`
			BytesIn         int    `json:"bytesIn"`
			Input           string `json:"input"`
			FlowFilesOut    int    `json:"flowFilesOut"`
			BytesOut        int    `json:"bytesOut"`
			Output          string `json:"output"`
			FlowFilesQueued int    `json:"flowFilesQueued"`
			BytesQueued     int    `json:"bytesQueued"`
			Queued          string `json:"queued"`
			QueuedSize      string `json:"queuedSize"`
			QueuedCount     string `json:"queuedCount"`
			PercentUseCount int    `json:"percentUseCount"`
			PercentUseBytes int    `json:"percentUseBytes"`
		} `json:"aggregateSnapshot"`
	} `json:"status"`
	Bends              []interface{} `json:"bends"`
	LabelIndex         int           `json:"labelIndex"`
	ZIndex             int           `json:"zIndex"`
	SourceID           string        `json:"sourceId"`
	SourceGroupID      string        `json:"sourceGroupId"`
	SourceType         string        `json:"sourceType"`
	DestinationID      string        `json:"destinationId"`
	DestinationGroupID string        `json:"destinationGroupId"`
	DestinationType    string        `json:"destinationType"`
}

type Flow struct {
	ProcessGroups       []ProcessGroup `json:"processGroups"`
	RemoteProcessGroups []interface{}  `json:"remoteProcessGroups"`
	Processors          []Processor    `json:"processors"`
	InputPorts          []interface{}  `json:"inputPorts"`
	OutputPorts         []interface{}  `json:"outputPorts"`
	Connections         []Connection   `json:"connections"`
	Labels              []struct {
		Revision struct {
			ClientID string `json:"clientId"`
			Version  int    `json:"version"`
		} `json:"revision"`
		ID       string `json:"id"`
		URI      string `json:"uri"`
		Position struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"position"`
		Permissions struct {
			CanRead  bool `json:"canRead"`
			CanWrite bool `json:"canWrite"`
		} `json:"permissions"`
		Dimensions struct {
			Width  float64 `json:"width"`
			Height float64 `json:"height"`
		} `json:"dimensions"`
		Component struct {
			ID            string `json:"id"`
			ParentGroupID string `json:"parentGroupId"`
			Position      struct {
				X float64 `json:"x"`
				Y float64 `json:"y"`
			} `json:"position"`
			Label  string  `json:"label"`
			Width  float64 `json:"width"`
			Height float64 `json:"height"`
			Style  struct {
				FontSize string `json:"font-size"`
			} `json:"style"`
		} `json:"component"`
	} `json:"labels"`
	Funnels []interface{} `json:"funnels"`
}

type ProcessGroupFlow struct {
	ID            string `json:"id"`
	URI           string `json:"uri"`
	ParentGroupID string `json:"parentGroupId"`
	Breadcrumb    struct {
		ID          string `json:"id"`
		Permissions struct {
			CanRead  bool `json:"canRead"`
			CanWrite bool `json:"canWrite"`
		} `json:"permissions"`
		Breadcrumb struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"breadcrumb"`
		ParentBreadcrumb struct {
			ID          string `json:"id"`
			Permissions struct {
				CanRead  bool `json:"canRead"`
				CanWrite bool `json:"canWrite"`
			} `json:"permissions"`
			Breadcrumb struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"breadcrumb"`
			ParentBreadcrumb struct {
				ID          string `json:"id"`
				Permissions struct {
					CanRead  bool `json:"canRead"`
					CanWrite bool `json:"canWrite"`
				} `json:"permissions"`
				Breadcrumb struct {
					ID   string `json:"id"`
					Name string `json:"name"`
				} `json:"breadcrumb"`
			} `json:"parentBreadcrumb"`
		} `json:"parentBreadcrumb"`
	} `json:"breadcrumb"`
	Flow          Flow   `json:"flow"`
	LastRefreshed string `json:"lastRefreshed"`
}

type ProcessGroupEntity struct {
	Permissions      Permissions      `json:"permissions"`
	ProcessGroupFlow ProcessGroupFlow `json:"processGroupFlow"`
}

type UpdateComponent struct {
	ID       string   `json:"id"`
	Position Position `json:"position"`
}

type UpdateProcessorEntity struct {
	Component                    UpdateComponent `json:"component"`
	DisconnectedNodeAcknowledged bool            `json:"disconnectedNodeAcknowledged"`
	Revision                     Revision        `json:"revision"`
}

type UpdateProcessorGroupEntity UpdateProcessorEntity

type InsertComponent struct {
	Bundle   Bundle   `json:"bundle"`
	Name     string   `json:"name"`
	Position Position `json:"position"`
	TypeID   string   `json:"type"`
}

type InsertProcessorEntity struct {
	Component                    InsertComponent `json:"component"`
	DisconnectedNodeAcknowledged bool            `json:"disconnectedNodeAcknowledged"`
	Revision                     Revision        `json:"revision"`
}

type InsertProcessGroupEntity struct {
	Component struct{
		Name     string   `json:"name"`
		Position Position `json:"position"`
	} `json:"component"`
	DisconnectedNodeAcknowledged bool            `json:"disconnectedNodeAcknowledged"`
	Revision                     Revision        `json:"revision"`
}

type CreatedSnippet struct {
	ID         string              `json:"id"`
	Processors map[string]Revision `json:"processors"`
}

type CreateSnippet struct {
	ParentGroupId string              `json:"parentGroupId"`
	Processors    map[string]Revision `json:"processors"`
	ProcessGroups map[string]Revision `json:"processGroups"`
}

type CreateSnippetEntity struct {
	DisconnectedNodeAcknowledged bool          `json:"disconnectedNodeAcknowledged"`
	Snippet                      CreateSnippet `json:"snippet"`
}

type CreatedSnippetEntity struct {
	Snippet CreatedSnippet `json:"snippet"`
}

type UpdateSnippetEntity struct {
	DisconnectedNodeAcknowledged bool          `json:"disconnectedNodeAcknowledged"`
	Snippet struct{
		ID string `json:"id"`
		ParentGroupId string `json:"parentGroupId"`
	} `json:"snippet"`
}