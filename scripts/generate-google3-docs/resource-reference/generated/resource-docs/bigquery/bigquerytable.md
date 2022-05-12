{# AUTOGENERATED. DO NOT EDIT. #}

{% extends "config-connector/_base.html" %}

{% block page_title %}BigQueryTable{% endblock %}
{% block body %}


<table>
<thead>
<tr>
<th><strong>Property</strong></th>
<th><strong>Value</strong></th>
</tr>
</thead>
<tbody>
<tr>
<td>{{gcp_name_short}} Service Name</td>
<td>BigQuery</td>
</tr>
<tr>
<td>{{gcp_name_short}} Service Documentation</td>
<td><a href="/bigquery/docs/">/bigquery/docs/</a></td>
</tr>
<tr>
<td>{{gcp_name_short}} REST Resource Name</td>
<td>bigquery/v2/tables</td>
</tr>
<tr>
<td>{{gcp_name_short}} Rest Resource Documentation</td>
<td><a href="/bigquery/docs/reference/rest/v2/tables">/bigquery/docs/reference/rest/v2/tables</a></td>
</tr>
<tr>
<td>{{product_name_short}} Resource Short Names</td>
<td>gcpbigquerytable<br>gcpbigquerytables<br>bigquerytable</td>
</tr>
<tr>
<td>{{product_name_short}} Service Name</td>
<td>bigquery.googleapis.com</td>
</tr>
<tr>
<td>{{product_name_short}} Resource Fully Qualified Name</td>
<td>bigquerytables.bigquery.cnrm.cloud.google.com</td>
</tr>

<tr>
    <td>Can Be Referenced by IAMPolicy/IAMPolicyMember</td>
    <td>Yes</td>
</tr>

<tr>
    <td>Supports IAM Conditions</td>
    <td>Yes</td>
</tr>

<tr>
    <td>Supports IAM Audit Configs</td>
    <td>No</td>
</tr>
<tr>
    <td>IAM External Reference Format</td>
    <td>
        
        <p>{% verbatim %}projects/{{project}}/datasets/{{dataset_id}}/tables/{{table_id}}{% endverbatim %}</p>
        
    </td>
</tr>


</tbody>
</table>

## Custom Resource Definition Properties


### Annotations
<table class="properties responsive">
<thead>
    <tr>
        <th colspan="2">Fields</th>
    </tr>
</thead>
<tbody>
    <tr>
        <td><code>cnrm.cloud.google.com/project-id</code></td>
    </tr>
</tbody>
</table>


### Spec
#### Schema
  ```yaml
  clustering:
  - string
  datasetRef:
    external: string
    name: string
    namespace: string
  description: string
  encryptionConfiguration:
    kmsKeyRef:
      external: string
      name: string
      namespace: string
    kmsKeyVersion: string
  expirationTime: integer
  externalDataConfiguration:
    autodetect: boolean
    compression: string
    csvOptions:
      allowJaggedRows: boolean
      allowQuotedNewlines: boolean
      encoding: string
      fieldDelimiter: string
      quote: string
      skipLeadingRows: integer
    googleSheetsOptions:
      range: string
      skipLeadingRows: integer
    hivePartitioningOptions:
      mode: string
      requirePartitionFilter: boolean
      sourceUriPrefix: string
    ignoreUnknownValues: boolean
    maxBadRecords: integer
    schema: string
    sourceFormat: string
    sourceUris:
    - string
  friendlyName: string
  materializedView:
    enableRefresh: boolean
    query: string
    refreshIntervalMs: integer
  rangePartitioning:
    field: string
    range:
      end: integer
      interval: integer
      start: integer
  resourceID: string
  schema: string
  timePartitioning:
    expirationMs: integer
    field: string
    requirePartitionFilter: boolean
    type: string
  view:
    query: string
    useLegacySql: boolean
  ```

<table class="properties responsive">
<thead>
    <tr>
        <th colspan="2">Fields</th>
    </tr>
</thead>
<tbody>
    <tr>
        <td>
            <p><code>clustering</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">list (string)</code></p>
            <p>{% verbatim %}Specifies column names to use for data clustering. Up to four top-level columns are allowed, and should be specified in descending priority order.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>clustering[]</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>datasetRef</code></p>
            <p><i>Required</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>datasetRef.external</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Allowed value: The `name` field of a `BigQueryDataset` resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>datasetRef.name</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>datasetRef.namespace</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>description</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The field description.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>encryptionConfiguration</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Immutable. Specifies how the table should be encrypted. If left blank, the table will be encrypted with a Google-managed key; that process is transparent to the user.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>encryptionConfiguration.kmsKeyRef</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>encryptionConfiguration.kmsKeyRef.external</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Allowed value: The `selfLink` field of a `KMSCryptoKey` resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>encryptionConfiguration.kmsKeyRef.name</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>encryptionConfiguration.kmsKeyRef.namespace</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>encryptionConfiguration.kmsKeyVersion</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The self link or full name of the kms key version used to encrypt this table.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>expirationTime</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}The time when this table expires, in milliseconds since the epoch. If not present, the table will persist indefinitely. Expired tables will be deleted and their storage reclaimed.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>externalDataConfiguration</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Describes the data format, location, and other properties of a table stored outside of BigQuery. By defining these properties, the data source can then be queried as if it were a standard BigQuery table.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>externalDataConfiguration.autodetect</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">boolean</code></p>
            <p>{% verbatim %}Let BigQuery try to autodetect the schema and format of the table.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>externalDataConfiguration.compression</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The compression type of the data source. Valid values are "NONE" or "GZIP".{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>externalDataConfiguration.csvOptions</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Additional properties to set if source_format is set to "CSV".{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>externalDataConfiguration.csvOptions.allowJaggedRows</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">boolean</code></p>
            <p>{% verbatim %}Indicates if BigQuery should accept rows that are missing trailing optional columns.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>externalDataConfiguration.csvOptions.allowQuotedNewlines</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">boolean</code></p>
            <p>{% verbatim %}Indicates if BigQuery should allow quoted data sections that contain newline characters in a CSV file. The default value is false.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>externalDataConfiguration.csvOptions.encoding</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The character encoding of the data. The supported values are UTF-8 or ISO-8859-1.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>externalDataConfiguration.csvOptions.fieldDelimiter</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The separator for fields in a CSV file.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>externalDataConfiguration.csvOptions.quote</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>externalDataConfiguration.csvOptions.skipLeadingRows</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}The number of rows at the top of a CSV file that BigQuery will skip when reading the data.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>externalDataConfiguration.googleSheetsOptions</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Additional options if source_format is set to "GOOGLE_SHEETS".{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>externalDataConfiguration.googleSheetsOptions.range</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Range of a sheet to query from. Only used when non-empty. At least one of range or skip_leading_rows must be set. Typical format: "sheet_name!top_left_cell_id:bottom_right_cell_id" For example: "sheet1!A1:B20".{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>externalDataConfiguration.googleSheetsOptions.skipLeadingRows</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}The number of rows at the top of the sheet that BigQuery will skip when reading the data. At least one of range or skip_leading_rows must be set.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>externalDataConfiguration.hivePartitioningOptions</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}When set, configures hive partitioning support. Not all storage formats support hive partitioning -- requesting hive partitioning on an unsupported format will lead to an error, as will providing an invalid specification.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>externalDataConfiguration.hivePartitioningOptions.mode</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}When set, what mode of hive partitioning to use when reading data.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>externalDataConfiguration.hivePartitioningOptions.requirePartitionFilter</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">boolean</code></p>
            <p>{% verbatim %}If set to true, queries over this table require a partition filter that can be used for partition elimination to be specified.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>externalDataConfiguration.hivePartitioningOptions.sourceUriPrefix</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}When hive partition detection is requested, a common for all source uris must be required. The prefix must end immediately before the partition key encoding begins.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>externalDataConfiguration.ignoreUnknownValues</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">boolean</code></p>
            <p>{% verbatim %}Indicates if BigQuery should allow extra values that are not represented in the table schema. If true, the extra values are ignored. If false, records with extra columns are treated as bad records, and if there are too many bad records, an invalid error is returned in the job result. The default value is false.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>externalDataConfiguration.maxBadRecords</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}The maximum number of bad records that BigQuery can ignore when reading data.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>externalDataConfiguration.schema</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Immutable. A JSON schema for the external table. Schema is required for CSV and JSON formats and is disallowed for Google Cloud Bigtable, Cloud Datastore backups, and Avro formats when using external tables.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>externalDataConfiguration.sourceFormat</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The data format. Supported values are: "CSV", "GOOGLE_SHEETS", "NEWLINE_DELIMITED_JSON", "AVRO", "PARQUET", "ORC" and "DATASTORE_BACKUP". To use "GOOGLE_SHEETS" the scopes must include "googleapis.com/auth/drive.readonly".{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>externalDataConfiguration.sourceUris</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">list (string)</code></p>
            <p>{% verbatim %}A list of the fully-qualified URIs that point to your data in Google Cloud.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>externalDataConfiguration.sourceUris[]</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>friendlyName</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}A descriptive name for the table.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>materializedView</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}If specified, configures this table as a materialized view.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>materializedView.enableRefresh</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">boolean</code></p>
            <p>{% verbatim %}Specifies if BigQuery should automatically refresh materialized view when the base table is updated. The default is true.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>materializedView.query</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Immutable. A query whose result is persisted.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>materializedView.refreshIntervalMs</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}Specifies maximum frequency at which this materialized view will be refreshed. The default is 1800000.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>rangePartitioning</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}If specified, configures range-based partitioning for this table.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>rangePartitioning.field</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Immutable. The field used to determine how to create a range-based partition.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>rangePartitioning.range</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Information required to partition based on ranges. Structure is documented below.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>rangePartitioning.range.end</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}End of the range partitioning, exclusive.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>rangePartitioning.range.interval</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}The width of each range within the partition.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>rangePartitioning.range.start</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}Start of the range partitioning, inclusive.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>resourceID</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Immutable. Optional. The tableId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>schema</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}A JSON schema for the table.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>timePartitioning</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}If specified, configures time-based partitioning for this table.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>timePartitioning.expirationMs</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}Number of milliseconds for which to keep the storage for a partition.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>timePartitioning.field</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Immutable. The field used to determine how to create a time-based partition. If time-based partitioning is enabled without this value, the table is partitioned based on the load time.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>timePartitioning.requirePartitionFilter</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">boolean</code></p>
            <p>{% verbatim %}If set to true, queries over this table require a partition filter that can be used for partition elimination to be specified.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>timePartitioning.type</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The supported types are DAY, HOUR, MONTH, and YEAR, which will generate one partition per day, hour, month, and year, respectively.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>view</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}If specified, configures this table as a view.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>view.query</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}A query that BigQuery executes when the view is referenced.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>view.useLegacySql</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">boolean</code></p>
            <p>{% verbatim %}Specifies whether to use BigQuery's legacy SQL for this view. The default value is true. If set to false, the view will use BigQuery's standard SQL.{% endverbatim %}</p>
        </td>
    </tr>
</tbody>
</table>


<p>{% verbatim %}* Field is required when parent field is specified{% endverbatim %}</p>


### Status
#### Schema
  ```yaml
  conditions:
  - lastTransitionTime: string
    message: string
    reason: string
    status: string
    type: string
  creationTime: integer
  etag: string
  lastModifiedTime: integer
  location: string
  numBytes: integer
  numLongTermBytes: integer
  numRows: integer
  observedGeneration: integer
  selfLink: string
  type: string
  ```

<table class="properties responsive">
<thead>
    <tr>
        <th colspan="2">Fields</th>
    </tr>
</thead>
<tbody>
    <tr>
        <td><code>conditions</code></td>
        <td>
            <p><code class="apitype">list (object)</code></p>
            <p>{% verbatim %}Conditions represent the latest available observation of the resource's current state.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[]</code></td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].lastTransitionTime</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Last time the condition transitioned from one status to another.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].message</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Human-readable message indicating details about last transition.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].reason</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Unique, one-word, CamelCase reason for the condition's last transition.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].status</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Status is the status of the condition. Can be True, False, Unknown.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].type</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Type is the type of the condition.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>creationTime</code></td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}The time when this table was created, in milliseconds since the epoch.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>etag</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}A hash of the resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>lastModifiedTime</code></td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}The time when this table was last modified, in milliseconds since the epoch.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>location</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The geographic location where the table resides. This value is inherited from the dataset.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>numBytes</code></td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}The geographic location where the table resides. This value is inherited from the dataset.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>numLongTermBytes</code></td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}The number of bytes in the table that are considered "long-term storage".{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>numRows</code></td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}The number of rows of data in this table, excluding any data in the streaming buffer.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>observedGeneration</code></td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>selfLink</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The URI of the created resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>type</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Describes the table type.{% endverbatim %}</p>
        </td>
    </tr>
</tbody>
</table>

## Sample YAML(s)

### Typical Use Case
  ```yaml
  # Copyright 2020 Google LLC
  #
  # Licensed under the Apache License, Version 2.0 (the "License");
  # you may not use this file except in compliance with the License.
  # You may obtain a copy of the License at
  #
  #     http://www.apache.org/licenses/LICENSE-2.0
  #
  # Unless required by applicable law or agreed to in writing, software
  # distributed under the License is distributed on an "AS IS" BASIS,
  # WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  # See the License for the specific language governing permissions and
  # limitations under the License.
  
  apiVersion: bigquery.cnrm.cloud.google.com/v1beta1
  kind: BigQueryTable
  metadata:
    name: bigquerytablesample
    labels:
      data-source: "external"
      schema-type: "auto-junk"
  spec:
    description: "BigQuery Sample Table"
    datasetRef:
      name: bigquerytabledep
    friendlyName: bigquerytable-sample
    externalDataConfiguration:
      autodetect: true
      compression: NONE
      ignoreUnknownValues: false
      maxBadRecords: 10
      sourceFormat: CSV
      sourceUris:
        - "gs://gcp-public-data-landsat/LC08/01/044/034/LC08_L1GT_044034_20130330_20170310_01_T2/LC08_L1GT_044034_20130330_20170310_01_T2_ANG.txt"
        - "gs://gcp-public-data-landsat/LC08/01/044/034/LC08_L1GT_044034_20130330_20180201_01_T2/LC08_L1GT_044034_20130330_20180201_01_T2_ANG.txt"
  ---
  apiVersion: bigquery.cnrm.cloud.google.com/v1beta1
  kind: BigQueryDataset
  metadata:
    name: bigquerytabledep
  spec:
    friendlyName: bigquerytable-dep
  ```


{% endblock %}