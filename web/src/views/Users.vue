<template lang="pug">
mixin IDColumn
  el-table-column(
    prop="id"
    key="id"
    label="ID"
    width="80"
    sortable
  )

mixin AccountColumn
  el-table-column(
    prop="account"
    key="account"
    label="账户"
    width="120"
  )

mixin StatusColumn
  el-table-column(
    label="状态"
    width="100"
  ): template(
    #default="scope"
  ) {{ getStatusDesc(scope.row.status) }}

mixin RolesColumn
  el-table-column(
    label="角色"
  ): template(
    #default="scope"
  ): ul
    li(
      v-for="role in scope.row.roles"
      :key="role"
    ) {{ role }}

mixin UpdatedAtColumn
  el-table-column(
    prop="updatedAt"
    key="updatedAt"
    label="更新于"
    width="160"
    sortable
  ): template(
    #default="scope"
  ): time-formater(
    :time="scope.row.updatedAt"
  )

mixin OpColumn
  el-table-column(
    label="操作"
    width="120"
  ): template(
    #default="scope"
  ): el-button.op(
    type="text"
    size="small"
    @click="modify(scope.row)"
  ) 编辑

mixin Pagination
  el-pagination.pagination(
    layout="prev, pager, next, sizes"
    :current-page="currentPage"
    :page-size="query.limit"
    :page-sizes="pageSizes"
    :total="users.count"
    @size-change="handleSizeChange"
    @current-change="handleCurrentChange"
  )

.users(
  v-loading="!inited"
)
  el-card(
    v-if="!editMode"
  )
    template(
      #header
    )
      i.el-icon-user-solid
      span 用户列表
    //- 筛选条件
    base-filter(
      :fields="filterFields"
      v-if="inited"
      :filter="filter"
    )
    div(
      v-loading="users.processing"
    ): el-table(
      :data="users.items"
      row-key="id"
      stripe
      @sort-change="handleSortChange"
    )
      //- 用户ID
      +IDColumn
      //- 用户账号
      +AccountColumn

      //- 用户状态
      +StatusColumn

      //- 用户角色
      +RolesColumn

      //- 更新时间
      +UpdatedAtColumn
      
      //- 操作
      +OpColumn

    //- 分页
    +Pagination
  //- 用户编辑
  user(
    v-else
  )
</template>

<script lang="ts">
import { defineComponent, onUnmounted } from "vue";

import useUserState, {
  userList,
  userListRole,
  userListClear,
} from "../states/user";
import useCommonState, { commonListStatus } from "../states/common";

import BaseFilter from "../components/base/Filter.vue";
import BaseTooltip from "../components/Tooltip.vue";
import TimeFormater from "../components/TimeFormater.vue";
import User from "../components/User.vue";
import { PAGE_SIZES } from "../constants/common";
import FilterTable from "../mixins/FilterTable";

const roleSelectList = [];
const statusSelectList = [];
const filterFields = [
  {
    label: "用户角色：",
    key: "role",
    type: "select",
    options: roleSelectList,
    span: 6,
  },
  {
    label: "用户状态：",
    key: "status",
    type: "select",
    options: statusSelectList,
    span: 6,
  },
  {
    label: "关键字：",
    key: "keyword",
    placeholder: "请输入关键字",
    clearable: true,
    span: 6,
  },
  {
    label: "",
    type: "filter",
    span: 6,
    labelWidth: "0px",
  },
];

export default defineComponent({
  name: "Users",
  components: {
    BaseFilter,
    BaseTooltip,
    TimeFormater,
    User,
  },
  mixins: [FilterTable],
  setup() {
    const commonState = useCommonState();
    const userState = useUserState();
    // 清理数据减少内存占用
    onUnmounted(() => {
      userListClear();
    });
    return {
      users: userState.users,
      userRoles: userState.roles,
      statuses: commonState.statuses,
      getStatusDesc: (status: number): string => {
        let desc = "";
        commonState.statuses.items.forEach((item) => {
          if (item.value === status) {
            desc = item.name;
          }
        });
        return desc;
      },
    };
  },
  data() {
    return {
      inited: false,
      filterFields: null,
      pageSizes: PAGE_SIZES,
      query: {
        offset: 0,
        limit: PAGE_SIZES[0],
        order: "-updatedAt",
      },
    };
  },
  async beforeMount() {
    try {
      await userListRole();
      await commonListStatus();

      // 重置
      roleSelectList.length = 0;
      roleSelectList.push({
        name: "所有",
        value: "",
      });
      roleSelectList.push(...this.userRoles.items);

      // 重置
      statusSelectList.length = 0;
      statusSelectList.push({
        name: "所有",
        value: "",
      });
      statusSelectList.push(...this.statuses.items);

      this.filterFields = filterFields;
    } catch (err) {
      this.$error(err);
    } finally {
      this.inited = true;
    }
  },
  methods: {
    async fetch() {
      const { users, query } = this;
      if (users.processing) {
        return;
      }
      try {
        await userList(query);
      } catch (err) {
        this.$error(err);
      }
    },
  },
});
</script>

<style lang="stylus" scoped>
@import "../common";
.users
  margin $mainMargin
  i
    margin-right 5px
  ul
    li
      list-style inside
.selector, .submit
  width 100%
.pagination
  text-align right
  margin-top 15px
</style>
