/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type Table struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TableSpec   `json:"spec,omitempty"`
	Status            TableStatus `json:"status,omitempty"`
}

type TableSpecEncryptionConfiguration struct {
	// The self link or full name of a key which should be used to encrypt this table. Note that the default bigquery service account will need to have encrypt/decrypt permissions on this key - you may want to see the google_bigquery_default_service_account datasource and the google_kms_crypto_key_iam_binding resource.
	KmsKeyName *string `json:"kmsKeyName" tf:"kms_key_name"`
}

type TableSpecExternalDataConfigurationCsvOptions struct {
	// Indicates if BigQuery should accept rows that are missing trailing optional columns.
	// +optional
	AllowJaggedRows *bool `json:"allowJaggedRows,omitempty" tf:"allow_jagged_rows"`
	// Indicates if BigQuery should allow quoted data sections that contain newline characters in a CSV file. The default value is false.
	// +optional
	AllowQuotedNewlines *bool `json:"allowQuotedNewlines,omitempty" tf:"allow_quoted_newlines"`
	// The character encoding of the data. The supported values are UTF-8 or ISO-8859-1.
	// +optional
	Encoding *string `json:"encoding,omitempty" tf:"encoding"`
	// The separator for fields in a CSV file.
	// +optional
	FieldDelimiter *string `json:"fieldDelimiter,omitempty" tf:"field_delimiter"`
	// The value that is used to quote data sections in a CSV file. If your data does not contain quoted sections, set the property value to an empty string. If your data contains quoted newline characters, you must also set the allow_quoted_newlines property to true. The API-side default is ", specified in Terraform escaped as \". Due to limitations with Terraform default values, this value is required to be explicitly set.
	Quote *string `json:"quote" tf:"quote"`
	// The number of rows at the top of a CSV file that BigQuery will skip when reading the data.
	// +optional
	SkipLeadingRows *int64 `json:"skipLeadingRows,omitempty" tf:"skip_leading_rows"`
}

type TableSpecExternalDataConfigurationGoogleSheetsOptions struct {
	// Range of a sheet to query from. Only used when non-empty. At least one of range or skip_leading_rows must be set. Typical format: "sheet_name!top_left_cell_id:bottom_right_cell_id" For example: "sheet1!A1:B20"
	// +optional
	Range *string `json:"range,omitempty" tf:"range"`
	// The number of rows at the top of the sheet that BigQuery will skip when reading the data. At least one of range or skip_leading_rows must be set.
	// +optional
	SkipLeadingRows *int64 `json:"skipLeadingRows,omitempty" tf:"skip_leading_rows"`
}

type TableSpecExternalDataConfigurationHivePartitioningOptions struct {
	// When set, what mode of hive partitioning to use when reading data.
	// +optional
	Mode *string `json:"mode,omitempty" tf:"mode"`
	// If set to true, queries over this table require a partition filter that can be used for partition elimination to be specified.
	// +optional
	RequirePartitionFilter *bool `json:"requirePartitionFilter,omitempty" tf:"require_partition_filter"`
	// When hive partition detection is requested, a common for all source uris must be required. The prefix must end immediately before the partition key encoding begins.
	// +optional
	SourceURIPrefix *string `json:"sourceURIPrefix,omitempty" tf:"source_uri_prefix"`
}

type TableSpecExternalDataConfiguration struct {
	// Let BigQuery try to autodetect the schema and format of the table.
	Autodetect *bool `json:"autodetect" tf:"autodetect"`
	// The compression type of the data source. Valid values are "NONE" or "GZIP".
	// +optional
	Compression *string `json:"compression,omitempty" tf:"compression"`
	// Additional properties to set if source_format is set to "CSV".
	// +optional
	CsvOptions *TableSpecExternalDataConfigurationCsvOptions `json:"csvOptions,omitempty" tf:"csv_options"`
	// Additional options if source_format is set to "GOOGLE_SHEETS".
	// +optional
	GoogleSheetsOptions *TableSpecExternalDataConfigurationGoogleSheetsOptions `json:"googleSheetsOptions,omitempty" tf:"google_sheets_options"`
	// When set, configures hive partitioning support. Not all storage formats support hive partitioning -- requesting hive partitioning on an unsupported format will lead to an error, as will providing an invalid specification.
	// +optional
	HivePartitioningOptions *TableSpecExternalDataConfigurationHivePartitioningOptions `json:"hivePartitioningOptions,omitempty" tf:"hive_partitioning_options"`
	// Indicates if BigQuery should allow extra values that are not represented in the table schema. If true, the extra values are ignored. If false, records with extra columns are treated as bad records, and if there are too many bad records, an invalid error is returned in the job result. The default value is false.
	// +optional
	IgnoreUnknownValues *bool `json:"ignoreUnknownValues,omitempty" tf:"ignore_unknown_values"`
	// The maximum number of bad records that BigQuery can ignore when reading data.
	// +optional
	MaxBadRecords *int64 `json:"maxBadRecords,omitempty" tf:"max_bad_records"`
	// A JSON schema for the external table. Schema is required for CSV and JSON formats and is disallowed for Google Cloud Bigtable, Cloud Datastore backups, and Avro formats when using external tables.
	// +optional
	Schema *string `json:"schema,omitempty" tf:"schema"`
	// The data format. Supported values are: "CSV", "GOOGLE_SHEETS", "NEWLINE_DELIMITED_JSON", "AVRO", "PARQUET", "ORC" and "DATASTORE_BACKUP". To use "GOOGLE_SHEETS" the scopes must include "googleapis.com/auth/drive.readonly".
	SourceFormat *string `json:"sourceFormat" tf:"source_format"`
	// A list of the fully-qualified URIs that point to your data in Google Cloud.
	SourceUris []string `json:"sourceUris" tf:"source_uris"`
}

type TableSpecMaterializedView struct {
	// Specifies if BigQuery should automatically refresh materialized view when the base table is updated. The default is true.
	// +optional
	EnableRefresh *bool `json:"enableRefresh,omitempty" tf:"enable_refresh"`
	// A query whose result is persisted.
	Query *string `json:"query" tf:"query"`
	// Specifies maximum frequency at which this materialized view will be refreshed. The default is 1800000
	// +optional
	RefreshIntervalMs *int64 `json:"refreshIntervalMs,omitempty" tf:"refresh_interval_ms"`
}

type TableSpecRangePartitioningRange struct {
	// End of the range partitioning, exclusive.
	End *int64 `json:"end" tf:"end"`
	// The width of each range within the partition.
	Interval *int64 `json:"interval" tf:"interval"`
	// Start of the range partitioning, inclusive.
	Start *int64 `json:"start" tf:"start"`
}

type TableSpecRangePartitioning struct {
	// The field used to determine how to create a range-based partition.
	Field *string `json:"field" tf:"field"`
	// Information required to partition based on ranges. Structure is documented below.
	Range *TableSpecRangePartitioningRange `json:"range" tf:"range"`
}

type TableSpecTimePartitioning struct {
	// Number of milliseconds for which to keep the storage for a partition.
	// +optional
	ExpirationMs *int64 `json:"expirationMs,omitempty" tf:"expiration_ms"`
	// The field used to determine how to create a time-based partition. If time-based partitioning is enabled without this value, the table is partitioned based on the load time.
	// +optional
	Field *string `json:"field,omitempty" tf:"field"`
	// If set to true, queries over this table require a partition filter that can be used for partition elimination to be specified.
	// +optional
	RequirePartitionFilter *bool `json:"requirePartitionFilter,omitempty" tf:"require_partition_filter"`
	// The supported types are DAY, HOUR, MONTH, and YEAR, which will generate one partition per day, hour, month, and year, respectively.
	Type *string `json:"type" tf:"type"`
}

type TableSpecView struct {
	// A query that BigQuery executes when the view is referenced.
	Query *string `json:"query" tf:"query"`
	// Specifies whether to use BigQuery's legacy SQL for this view. The default value is true. If set to false, the view will use BigQuery's standard SQL
	// +optional
	UseLegacySQL *bool `json:"useLegacySQL,omitempty" tf:"use_legacy_sql"`
}

type TableSpec struct {
	KubeformOutput *TableSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource TableSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type TableSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies column names to use for data clustering. Up to four top-level columns are allowed, and should be specified in descending priority order.
	// +optional
	// +kubebuilder:validation:MaxItems=4
	Clustering []string `json:"clustering,omitempty" tf:"clustering"`
	// The time when this table was created, in milliseconds since the epoch.
	// +optional
	CreationTime *int64 `json:"creationTime,omitempty" tf:"creation_time"`
	// The dataset ID to create the table in. Changing this forces a new resource to be created.
	DatasetID *string `json:"datasetID" tf:"dataset_id"`
	// Whether or not to allow Terraform to destroy the instance. Unless this field is set to false in Terraform state, a terraform destroy or terraform apply that would delete the instance will fail.
	// +optional
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection"`
	// The field description.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// Specifies how the table should be encrypted. If left blank, the table will be encrypted with a Google-managed key; that process is transparent to the user.
	// +optional
	EncryptionConfiguration *TableSpecEncryptionConfiguration `json:"encryptionConfiguration,omitempty" tf:"encryption_configuration"`
	// A hash of the resource.
	// +optional
	Etag *string `json:"etag,omitempty" tf:"etag"`
	// The time when this table expires, in milliseconds since the epoch. If not present, the table will persist indefinitely. Expired tables will be deleted and their storage reclaimed.
	// +optional
	ExpirationTime *int64 `json:"expirationTime,omitempty" tf:"expiration_time"`
	// Describes the data format, location, and other properties of a table stored outside of BigQuery. By defining these properties, the data source can then be queried as if it were a standard BigQuery table.
	// +optional
	ExternalDataConfiguration *TableSpecExternalDataConfiguration `json:"externalDataConfiguration,omitempty" tf:"external_data_configuration"`
	// A descriptive name for the table.
	// +optional
	FriendlyName *string `json:"friendlyName,omitempty" tf:"friendly_name"`
	// A mapping of labels to assign to the resource.
	// +optional
	Labels *map[string]string `json:"labels,omitempty" tf:"labels"`
	// The time when this table was last modified, in milliseconds since the epoch.
	// +optional
	LastModifiedTime *int64 `json:"lastModifiedTime,omitempty" tf:"last_modified_time"`
	// The geographic location where the table resides. This value is inherited from the dataset.
	// +optional
	Location *string `json:"location,omitempty" tf:"location"`
	// If specified, configures this table as a materialized view.
	// +optional
	MaterializedView *TableSpecMaterializedView `json:"materializedView,omitempty" tf:"materialized_view"`
	// The geographic location where the table resides. This value is inherited from the dataset.
	// +optional
	NumBytes *int64 `json:"numBytes,omitempty" tf:"num_bytes"`
	// The number of bytes in the table that are considered "long-term storage".
	// +optional
	NumLongTermBytes *int64 `json:"numLongTermBytes,omitempty" tf:"num_long_term_bytes"`
	// The number of rows of data in this table, excluding any data in the streaming buffer.
	// +optional
	NumRows *int64 `json:"numRows,omitempty" tf:"num_rows"`
	// The ID of the project in which the resource belongs.
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// If specified, configures range-based partitioning for this table.
	// +optional
	RangePartitioning *TableSpecRangePartitioning `json:"rangePartitioning,omitempty" tf:"range_partitioning"`
	// A JSON schema for the table.
	// +optional
	Schema *string `json:"schema,omitempty" tf:"schema"`
	// The URI of the created resource.
	// +optional
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link"`
	// A unique ID for the resource. Changing this forces a new resource to be created.
	TableID *string `json:"tableID" tf:"table_id"`
	// If specified, configures time-based partitioning for this table.
	// +optional
	TimePartitioning *TableSpecTimePartitioning `json:"timePartitioning,omitempty" tf:"time_partitioning"`
	// Describes the table type.
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
	// If specified, configures this table as a view.
	// +optional
	View *TableSpecView `json:"view,omitempty" tf:"view"`
}

type TableStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase status.Status `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// TableList is a list of Tables
type TableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Table CRD objects
	Items []Table `json:"items,omitempty"`
}
