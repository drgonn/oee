export interface TableListItem {
  id?: number;
  start_time: Date;
  end_time: Date;
  seconds: number;
  volt: number;
  amount: number;
  good: number;
  glue: number;
}

export interface TablePutItem {
  id?: number;
  start_time?: Date;
  end_time?: Date;
  seconds?: number;
  volt?: number;
  amount?: number;
  good?: number;
  glue?: number;
}

export interface TableListPagination {
  total: number;
  pageSize: number;
  current: number;
}

export interface TableListData {
  list: TableListItem[];
  pagination: Partial<TableListPagination>;
}

export interface TableListParams {
  status?: string;
  name?: string;
  desc?: string;
  key?: number;
  pageSize?: number;
  currentPage?: number;
  filter?: { [key: string]: any[] };
  sorter?: { [key: string]: any };
}
