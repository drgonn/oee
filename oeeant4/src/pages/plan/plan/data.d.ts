export interface TableListItem {
  id?: number;
  name: string;
  current: string;
  follow: string;
  week: number;
  time: Date;
}

export interface TablePutItem {
  id?: number;
  name?: string;
  week?: number;
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
