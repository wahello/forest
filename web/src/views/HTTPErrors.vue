<template lang="pug">
mixin AccountColumn
  el-table-column(
    prop="account"
    key="account"
    label="账户"
    width="120"
    fixed="left"
  )

mixin MethodColumn
  el-table-column(
    prop="method"
    key="method"
    label="Method"
    width="80"
  )

mixin RouteColumn
  el-table-column(
    prop="route"
    key="route"
    label="路由"
    width="180"
  )

mixin CategoryColumn
  el-table-column(
    prop="category"
    key="category"
    label="类型"
    width="200"
  )

mixin StatusColumn
  el-table-column(
    prop="status"
    key="status"
    label="状态码"
    width="80"
  )

mixin ExceptionColumn
  el-table-column(
    prop="exception"
    key="exception"
    label="异常"
    width="80"
  ): template(
    #default="scope"
  ): span {{scope.row.exception? "是" : "否"}}

mixin SessionColumn
  el-table-column(
    label="Session ID"
    :filters="sessionIDFilters"
    :filter-method="filterSession"
    width="110"
  ): template(
    #default="scope"
  ): base-tooltip(
    :content="scope.row.sid"
  )

mixin TrackIDColumn
  el-table-column(
    label="Track ID"
    :filters="trackIDFilters"
    :filter-method="filterTrack"
    width="90"
  ): template(
    #default="scope"
  ): base-tooltip(
    :content="scope.row.tid"
  )

mixin IPColumn
  el-table-column(
    prop="ip"
    key="ip"
    label="IP"
    width="100"
  )

mixin URIColumn
  el-table-column(
    label="URI"
    width="80"
  ): template(
    #default="scope"
  ): base-tooltip(
    icon="el-icon-info"
    :content="scope.row.uri"
  )

mixin ErrorColumn
  el-table-column(
    prop="error"
    key="error"
    label="Error"
    width="300"
  ): template(
    #default="scope"
  )
    p(
      v-if="scope.row.error"
      v-html=`scope.row.error.split(",").map(item => item.trim()).join("<br />")`
    )
    span(
      v-else
    ) --

mixin TimeColumn
  el-table-column(
    label="时间"
    prop="_time"
    key="_time"
    width="160"
    fixed="right"
  ): template(
    #default="scope"
  ): time-formater(
    :time="scope.row._time"
  )

el-card.httpErrors
  template(
    #header
  )
    i.el-icon-user-solid
    span HTTP出错查询 
  div(
    v-loading="httpErrors.processing"
  )
    base-filter(
      v-if="inited"
      :fields="filterFields"
      :filter="filter"
    )
    el-table(
      :data="httpErrors.items"
      row-key="_time"
      stripe
      :default-sort="{ prop: '_time', order: 'descending' }"
    )
      //- 账号
      +AccountColumn

      //- http method
      +MethodColumn

      //- http route
      +RouteColumn

      //- 分类
      +CategoryColumn

      //- 状态
      +StatusColumn

      //- 是否异常
      +ExceptionColumn

      //- session id
      +SessionColumn

      //- track id
      +TrackIDColumn
 
      //- ip
      +IPColumn
      
      //- uri 
      +URIColumn

      //- error 
      +ErrorColumn      

      //- 时间
      +TimeColumn


</template>

<script lang="ts">
import { defineComponent, onUnmounted } from "vue";

import { getDateTimeShortcuts, formatDateWithTZ } from "../helpers/util";
import BaseFilter from "../components/base/Filter.vue";
import BaseTooltip from "../components/Tooltip.vue";
import TimeFormater from "../components/TimeFormater.vue";
import BaseJson from "../components/base/JSON.vue";
import { PAGE_SIZES } from "../constants/common";
import FilterTable from "../mixins/FilterTable";
import useFluxState, {
  fluxListHTTPCategory,
  fluxListHTTPError,
  fluxListHTTPErrorClear,
} from "../states/flux";

// 最近一小时
const defaultDateRange = [new Date(Date.now() - 60 * 60 * 1000), new Date()];
const categories = [];
const filterFields = [
  {
    label: "账号：",
    key: "account",
    placeholder: "请输入要查询的账号",
    clearable: true,
    span: 6,
  },
  {
    label: "数量：",
    key: "limit",
    type: "number",
    placeholder: "请输入最大数量",
    clearable: true,
    defaultValue: 100,
    span: 6,
  },
  {
    label: "类型：",
    key: "category",
    type: "select",
    placeholder: "请选择出错类型",
    options: categories,
    span: 6,
  },
  {
    label: "异常：",
    key: "exception",
    type: "select",
    placeholder: "请选择是否异常出错",
    options: [
      {
        name: "全部",
        value: "",
      },
      {
        name: "是",
        value: "1",
      },
      {
        name: "否",
        value: "0",
      },
    ],
    span: 6,
  },
  {
    label: "时间：",
    key: "dateRange",
    type: "dateTimeRange",
    placeholder: ["开始日期", "结束日期"],
    shortcuts: getDateTimeShortcuts(["1h", "2h", "3h", "12h", "1d"]),
    defaultValue: defaultDateRange,
    span: 16,
  },
  {
    label: "",
    type: "filter",
    labelWidth: "0px",
    span: 8,
  },
];

function getUniqueKey(data: Record<string, unknown>[], key: string) {
  if (!data || !data.length) {
    return [];
  }
  const keys = {};
  data.forEach((item) => {
    if (item[key]) {
      keys[`${item[key]}`] = true;
    }
  });
  return Object.keys(keys).map((item) => {
    return {
      text: item,
      value: item,
    };
  });
}

export default defineComponent({
  name: "Trackers",
  components: {
    BaseFilter,
    BaseTooltip,
    TimeFormater,
    BaseJson,
  },
  mixins: [FilterTable],
  setup() {
    onUnmounted(() => {
      fluxListHTTPErrorClear();
    });
    const fluxState = useFluxState();
    return {
      httpErrors: fluxState.httpErrors,
      httpErrorCategories: fluxState.httpErrorCategories,
    };
  },
  data() {
    return {
      inited: false,
      disableBeforeMountFetch: true,
      filterFields,
      pageSizes: PAGE_SIZES,
      query: {
        dateRange: defaultDateRange,
        offset: 0,
        limit: 100,
        account: "",
        exception: false,
      },
    };
  },
  computed: {
    trackIDFilters() {
      return getUniqueKey(this.httpErrors.items, "tid");
    },
    sessionIDFilters() {
      return getUniqueKey(this.httpErrors.items, "sid");
    },
  },
  async beforeMount() {
    try {
      await fluxListHTTPCategory();
      categories.length = 0;
      categories.push({
        name: "全部",
        value: "",
      });
      this.httpErrorCategories.items.forEach((element) => {
        categories.push({
          name: element,
          value: element,
        });
      });
      this.inited = true;
    } catch (err) {
      this.$error(err);
    }
  },
  methods: {
    filterTrack(value, row) {
      return row.tid == value;
    },
    filterSession(value, row) {
      return row.sid == value;
    },
    async fetch() {
      const { httpErrors, query } = this;
      if (httpErrors.processing) {
        return;
      }
      const params = Object.assign({}, query);
      const value = params.dateRange;
      if (!value || value.length !== 2) {
        this.$erro("时间区间不能为空");
        return;
      }
      params.begin = formatDateWithTZ(value[0]);
      params.end = formatDateWithTZ(value[1]);
      delete params.dateRange;
      try {
        await fluxListHTTPError(params);
      } catch (err) {
        this.$error(err);
      }
    },
  },
});
</script>

<style lang="stylus" scoped>
@import "../common";
.httpErrors
  margin $mainMargin
  i
    margin-right 5px
.pagination
  text-align right
  margin-top 15px
</style>
