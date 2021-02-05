import { DownOutlined, PlusOutlined, QuestionCircleOutlined} from '@ant-design/icons';
import { Button, Divider, Dropdown, Menu, message, Input, Form, Modal, Tooltip, Select, InputNumber ,Upload,DatePicker  } from 'antd';
import React, { useState, useRef, useEffect} from 'react';
import { PageHeaderWrapper } from '@ant-design/pro-layout';
import ProTable, { ProColumns, ActionType } from '@ant-design/pro-table';

import {queryDeviceList} from "@/pages/oee/device/service";
import {queryBugtypeList} from "@/pages/oee/bugtype/service";
import { TableListItem } from './data.d';
import { queryBugList, updateBug , addBug , removeBug  } from './service';

const { TextArea } = Input;
const handleRemove = async (selectedRows: TableListItem[]) => {
  const hide = message.loading('正在删除');
  if (!selectedRows) return true;
  try {
    const res = await removeBug({
      ids: selectedRows.map((row) => row.id),
    });
    if (res.success){
    hide();
    message.success('删除成功，即将刷新');
    return true;
        }else{
          message.error(res.errmsg||'请求失败请重试！');
          hide();
          return;
        }
  } catch (error) {
    hide();
    message.error('删除失败，请重试');
    return false;
  }
};

const TableList: React.FC<{}> = () => {
  const [createModalVisible, handleModalVisible] = useState<boolean>(false);
  const [Devicelist, setDevicelist] = useState({ data: [] });
  const [Bugtypelist, setBugtypelist] = useState({ data: [] });
  const [updateModalVisible, handleUpdateModalVisible] = useState<boolean>(false);
  const [stepFormValues, setStepFormValues] = useState({});
  const [id,setId]=useState(0)
  const [formvalues,setValues]=useState({})
  const [form] = Form.useForm()
  const actionRef = useRef<ActionType>();
  const columns: ProColumns<TableListItem>[] = [
    {
      title: '设备名',
      dataIndex: 'device_name',
      valueType: 'text',
      renderFormItem:()=>{
        return(
            <Form.Item
              label=''
              name="device_id"
            >
              <Select
                placeholder="请选择设备..."
                onPopupScroll={handlePopupScrollDevice}
                allowClear
                showSearch
                optionFilterProp="children"
              >
                {Devicelist.data.length &&
                  Devicelist.data.map((obj) => {
                    return <Select.Option value={obj.id} key={obj.id}>{obj.name}</Select.Option>;
                  })}
              </Select>
            </Form.Item>
          )}
    },
    {
      title: '类型名',
      dataIndex: 'bugtype_name',
      valueType: 'text',
      renderFormItem:()=>{
        return(
            <Form.Item
              label=''
              name="bugtype_id"
            >
              <Select
                placeholder="请选择类型..."
                onPopupScroll={handlePopupScrollBugtype}
                allowClear
                showSearch
                optionFilterProp="children"
              >
                {Bugtypelist.data.length &&
                  Bugtypelist.data.map((obj) => {
                    return <Select.Option value={obj.id} key={obj.id}>{obj.name}</Select.Option>;
                  })}
              </Select>
            </Form.Item>
          )}
    },
    {
      title: '原因',
      dataIndex: 'reason',
      valueType: 'text',
      search: true,
    },
    {
      title: '时间',
      dataIndex: 'start_time',
      valueType: 'dateTime',
      search: false,
    },
    {
      title: '修好时间',
      dataIndex: 'end_time',
      valueType: 'dateTime',
      search: false,
    },
    {
      title: '操作',
      dataIndex: 'option',
      valueType: 'option',
      render: (_, record) => (
        <>
          <a
            onClick={() => {
              handleUpdateModalVisible(true);
              setValues(record);
              setId(record.id);
            }}
          >
            修改
          </a>
          <Divider type="vertical" />
        </>
      ),
    },
  ];

  const getDevicelist = async (obj) => {
    const Devicedata = await queryDeviceList(obj);
    if (Devicedata.success) {
      const {data} = Devicelist
      Devicedata.data = data.concat(Devicedata.data)
      setDevicelist(Devicedata);
    }
  };
useEffect(()=>{getDevicelist({pageindex:1})},[])
  const handlePopupScrollDevice = async (e) => {
    const { current, pagecount } = Devicelist;
    e.persist();
    const { target } = e;
    if (
      Math.ceil(target.scrollTop + target.offsetHeight) === target.scrollHeight &&
      current < pagecount
    ) {
      getDevicelist({ current: current + 1 });
    }
  };
  const getBugtypelist = async (obj) => {
    const Bugtypedata = await queryBugtypeList(obj);
    if (Bugtypedata.success) {
      const {data} = Bugtypelist
      Bugtypedata.data = data.concat(Bugtypedata.data)
      setBugtypelist(Bugtypedata);
    }
  };
useEffect(()=>{getBugtypelist({pageindex:1})},[])
  const handlePopupScrollBugtype = async (e) => {
    const { current, pagecount } = Bugtypelist;
    e.persist();
    const { target } = e;
    if (
      Math.ceil(target.scrollTop + target.offsetHeight) === target.scrollHeight &&
      current < pagecount
    ) {
      getBugtypelist({ current: current + 1 });
    }
  };
  const handleUpdate = ()=>{
    const hide=message.loading('正在提交...')
    let v_time = new Object();
    form
      .validateFields().then(async(values)=>{
      if (values.start_timeu) {
        v_time.start_time = values.start_timeu.format('YYYYMMDDHHmmss')
      }
      if (values.end_timeu) {
        v_time.end_time = values.end_timeu.format('YYYYMMDDHHmmss')
      }
      try{
        values.id = id
        const res=await updateBug({...values,...v_time
})
        if(res.success){
          hide()
          message.success('修改成功！')
          handleUpdateModalVisible(false);
          actionRef.current.reload();
        }else{
          message.error(res.errmsg||'请求失败请重试！');
          hide();
          return;
        }
      }catch(error){
        message.error('请求失败请重试！');
        hide();
      }
    })
  }
  const handleAdd = ()=>{
    let v_time = new Object();
    form
      .validateFields().then(async(values)=>{
      if (values.start_time) {
        v_time.start_time = values.start_time.format('YYYYMMDDHHmmss')
      }
      if (values.end_time) {
        v_time.end_time = values.end_time.format('YYYYMMDDHHmmss')
      }
      const hide=message.loading('正在提交...')
      try{
        const res=await addBug({...values,...v_time
})
        if(res.success){
          hide()
          message.success('创建成功！')
          handleModalVisible(false);
          actionRef.current.reload();
        }else{
          message.error(res.errmsg||'请求失败请重试！');
          hide();
          return;
        }
      }catch(error){
        message.error('请求失败请重试！');
        hide();
      }
    })
  }
  return (
    <PageHeaderWrapper>
      <ProTable<TableListItem>
        rowKey="id"
        actionRef={actionRef}
        rowSelection={{}}
        toolBarRender={(action, { selectedRows }) => [
          <Button type="primary" onClick={() => handleModalVisible(true)}>
            <PlusOutlined /> 新建
          </Button>,
          selectedRows && selectedRows.length > 0 && (
            <Dropdown
              overlay={
                <Menu
                  onClick={async (e) => {
                    if (e.key === 'remove') {
                      await handleRemove(selectedRows);
                      action.reload();
                    }
                  }}
                  selectedKeys={[]}
                >
                  <Menu.Item key="remove">批量删除</Menu.Item>
                </Menu>
              }
            >
              <Button>
                批量操作 <DownOutlined />
              </Button>
            </Dropdown>
          ),
        ]}
        request={(params, sorter, filter) => queryBugList({ ...params, sorter, filter })}
        columns={columns}
      />
      <Modal
        title="新建故障"
        visible={createModalVisible}
        onOk={handleAdd}
        onCancel={() => handleModalVisible(false)} 
      >
        <Form form={form} initialValues={formvalues}>
            <Form.Item
              label={
                <span>
                  选择设备 &nbsp;
                  <Tooltip title="选择需要的设备">
                    <QuestionCircleOutlined />
                  </Tooltip>
                </span>
              }
              name="device_id"
              rules={[
                {
                  required: true,
                  message: '请选择设备!',
                },
              ]}
            >
              <Select
                placeholder="请选择设备..."
                onPopupScroll={handlePopupScrollDevice}
                allowClear
                showSearch
                optionFilterProp="children"
              >
                {Devicelist.data.length &&
                  Devicelist.data.map((obj) => {
                    return <Select.Option value={obj.id} key={obj.id}>{obj.name}</Select.Option>;
                  })}
              </Select>
            </Form.Item>
            <Form.Item
              label={
                <span>
                  选择类型 &nbsp;
                  <Tooltip title="选择需要的类型">
                    <QuestionCircleOutlined />
                  </Tooltip>
                </span>
              }
              name="bugtype_id"
              rules={[
                {
                  required: true,
                  message: '请选择类型!',
                },
              ]}
            >
              <Select
                placeholder="请选择类型..."
                onPopupScroll={handlePopupScrollBugtype}
                allowClear
                showSearch
                optionFilterProp="children"
              >
                {Bugtypelist.data.length &&
                  Bugtypelist.data.map((obj) => {
                    return <Select.Option value={obj.id} key={obj.id}>{obj.name}</Select.Option>;
                  })}
              </Select>
            </Form.Item>
          <Form.Item
            name='reason'
            rules= {[{ required: false, message: '请输入名称!' }]}
            label="故障原因"
          >
            <TextArea rows={4} />
          </Form.Item>
          <Form.Item
            name='start_time'
            rules= {[{ required: true, message: '请输入名称!' }]}
            label="故障时间"
          >
            <DatePicker showTime/>
          </Form.Item>
          <Form.Item
            name='end_time'
            rules= {[{ required: false, message: '请输入名称!' }]}
            label="故障修好时间"
          >
            <DatePicker showTime/>
          </Form.Item>
        </Form>
      </Modal>
      <Modal
        title="编辑故障"
        visible={updateModalVisible}
        onOk={handleUpdate}
        onCancel={()=>{handleUpdateModalVisible(false)}}
      >
        <Form form={form} initialValues={formvalues}>
            <Form.Item
              label={
                <span>
                  选择设备 &nbsp;
                  <Tooltip title="选择需要的设备">
                    <QuestionCircleOutlined />
                  </Tooltip>
                </span>
              }
              name="device_id"
              rules={[
                {
                  required: true,
                  message: '请选择设备!',
                },
              ]}
            >
              <Select
                placeholder="请选择设备..."
                onPopupScroll={handlePopupScrollDevice}
                allowClear
                showSearch
                optionFilterProp="children"
              >
                {Devicelist.data.length &&
                  Devicelist.data.map((obj) => {
                    return <Select.Option value={obj.id} key={obj.id}>{obj.name}</Select.Option>;
                  })}
              </Select>
            </Form.Item>
            <Form.Item
              label={
                <span>
                  选择类型 &nbsp;
                  <Tooltip title="选择需要的类型">
                    <QuestionCircleOutlined />
                  </Tooltip>
                </span>
              }
              name="bugtype_id"
              rules={[
                {
                  required: true,
                  message: '请选择类型!',
                },
              ]}
            >
              <Select
                placeholder="请选择类型..."
                onPopupScroll={handlePopupScrollBugtype}
                allowClear
                showSearch
                optionFilterProp="children"
              >
                {Bugtypelist.data.length &&
                  Bugtypelist.data.map((obj) => {
                    return <Select.Option value={obj.id} key={obj.id}>{obj.name}</Select.Option>;
                  })}
              </Select>
            </Form.Item>
          <Form.Item
            name='reason'
            rules= {[{ required: false, message: '请输入名称!' }]}
            label="故障原因"
          >
            <TextArea rows={4} />
          </Form.Item>
          <Form.Item
            name='start_timeu'
            rules= {[{ required: true, message: '请输入名称!' }]}
            label="故障时间"
          >
            <DatePicker showTime/>
          </Form.Item>
          <Form.Item
            name='end_timeu'
            rules= {[{ required: false, message: '请输入名称!' }]}
            label="故障修好时间"
          >
            <DatePicker showTime/>
          </Form.Item>
        </Form>
      </Modal>
    </PageHeaderWrapper>
  );
};

export default TableList;
