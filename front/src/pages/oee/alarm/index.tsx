import { DownOutlined, PlusOutlined, QuestionCircleOutlined} from '@ant-design/icons';
import { Button, Divider, Dropdown, Menu, message, Input, Form, Modal, Tooltip, Select, InputNumber ,Upload,DatePicker  } from 'antd';
import React, { useState, useRef, useEffect} from 'react';
import { PageHeaderWrapper } from '@ant-design/pro-layout';
import ProTable, { ProColumns, ActionType } from '@ant-design/pro-table';

import {queryAlarmtypeList} from "@/pages/oee/alarmtype/service";
import {queryValveList} from "@/pages/oee/valve/service";
import {queryDeviceList} from "@/pages/oee/device/service";
import { TableListItem } from './data.d';
import { queryAlarmList, updateAlarm , addAlarm , removeAlarm  } from './service';

const { TextArea } = Input;
const handleRemove = async (selectedRows: TableListItem[]) => {
  const hide = message.loading('正在删除');
  if (!selectedRows) return true;
  try {
    const res = await removeAlarm({
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
  const [Alarmtypelist, setAlarmtypelist] = useState({ data: [] });
  const [Valvelist, setValvelist] = useState({ data: [] });
  const [Devicelist, setDevicelist] = useState({ data: [] });
  const [updateModalVisible, handleUpdateModalVisible] = useState<boolean>(false);
  const [stepFormValues, setStepFormValues] = useState({});
  const [id,setId]=useState(0)
  const [formvalues,setValues]=useState({})
  const [form] = Form.useForm()
  const actionRef = useRef<ActionType>();
  const columns: ProColumns<TableListItem>[] = [
    {
      title: '代码',
      dataIndex: 'alarmtype_code',
      valueType: 'text',
      renderFormItem:()=>{
        return(
            <Form.Item
              label=''
              name="alarmtype_id"
            >
              <Select
                placeholder="请选择阀..."
                onPopupScroll={handlePopupScrollAlarmtype}
                allowClear
                showSearch
                optionFilterProp="children"
              >
                {Alarmtypelist.data.length &&
                  Alarmtypelist.data.map((obj) => {
                    return <Select.Option value={obj.id} key={obj.id}>{obj.name}</Select.Option>;
                  })}
              </Select>
            </Form.Item>
          )}
    },
    {
      title: '含义',
      dataIndex: 'alarmtype_mean',
      valueType: 'text',
      renderFormItem:()=>{
        return(
            <Form.Item
              label=''
              name="alarmtype_id"
            >
              <Select
                placeholder="请选择阀..."
                onPopupScroll={handlePopupScrollAlarmtype}
                allowClear
                showSearch
                optionFilterProp="children"
              >
                {Alarmtypelist.data.length &&
                  Alarmtypelist.data.map((obj) => {
                    return <Select.Option value={obj.id} key={obj.id}>{obj.name}</Select.Option>;
                  })}
              </Select>
            </Form.Item>
          )}
    },
    {
      title: '可能原因',
      dataIndex: 'alarmtype_cause',
      valueType: 'text',
      renderFormItem:()=>{
        return(
            <Form.Item
              label=''
              name="alarmtype_id"
            >
              <Select
                placeholder="请选择阀..."
                onPopupScroll={handlePopupScrollAlarmtype}
                allowClear
                showSearch
                optionFilterProp="children"
              >
                {Alarmtypelist.data.length &&
                  Alarmtypelist.data.map((obj) => {
                    return <Select.Option value={obj.id} key={obj.id}>{obj.name}</Select.Option>;
                  })}
              </Select>
            </Form.Item>
          )}
    },
    {
      title: '解决方案',
      dataIndex: 'alarmtype_solution',
      valueType: 'text',
      renderFormItem:()=>{
        return(
            <Form.Item
              label=''
              name="alarmtype_id"
            >
              <Select
                placeholder="请选择阀..."
                onPopupScroll={handlePopupScrollAlarmtype}
                allowClear
                showSearch
                optionFilterProp="children"
              >
                {Alarmtypelist.data.length &&
                  Alarmtypelist.data.map((obj) => {
                    return <Select.Option value={obj.id} key={obj.id}>{obj.name}</Select.Option>;
                  })}
              </Select>
            </Form.Item>
          )}
    },
    {
      title: '阀名',
      dataIndex: 'valve_name',
      valueType: 'text',
      renderFormItem:()=>{
        return(
            <Form.Item
              label=''
              name="valve_id"
            >
              <Select
                placeholder="请选择阀..."
                onPopupScroll={handlePopupScrollValve}
                allowClear
                showSearch
                optionFilterProp="children"
              >
                {Valvelist.data.length &&
                  Valvelist.data.map((obj) => {
                    return <Select.Option value={obj.id} key={obj.id}>{obj.name}</Select.Option>;
                  })}
              </Select>
            </Form.Item>
          )}
    },
    {
      title: 'id',
      dataIndex: 'valve_id',
      valueType: 'digit',
      renderFormItem:()=>{
        return(
            <Form.Item
              label=''
              name="valve_id"
            >
              <Select
                placeholder="请选择阀..."
                onPopupScroll={handlePopupScrollValve}
                allowClear
                showSearch
                optionFilterProp="children"
              >
                {Valvelist.data.length &&
                  Valvelist.data.map((obj) => {
                    return <Select.Option value={obj.id} key={obj.id}>{obj.name}</Select.Option>;
                  })}
              </Select>
            </Form.Item>
          )}
    },
    {
      title: '时间',
      dataIndex: 'create_time',
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

  const getAlarmtypelist = async (obj) => {
    const Alarmtypedata = await queryAlarmtypeList(obj);
    if (Alarmtypedata.success) {
      const {data} = Alarmtypelist
      Alarmtypedata.data = data.concat(Alarmtypedata.data)
      setAlarmtypelist(Alarmtypedata);
    }
  };
useEffect(()=>{getAlarmtypelist({pageindex:1})},[])
  const handlePopupScrollAlarmtype = async (e) => {
    const { current, pagecount } = Alarmtypelist;
    e.persist();
    const { target } = e;
    if (
      Math.ceil(target.scrollTop + target.offsetHeight) === target.scrollHeight &&
      current < pagecount
    ) {
      getAlarmtypelist({ current: current + 1 });
    }
  };
  const getValvelist = async (obj) => {
    const Valvedata = await queryValveList(obj);
    if (Valvedata.success) {
      const {data} = Valvelist
      Valvedata.data = data.concat(Valvedata.data)
      setValvelist(Valvedata);
    }
  };
useEffect(()=>{getValvelist({pageindex:1})},[])
  const handlePopupScrollValve = async (e) => {
    const { current, pagecount } = Valvelist;
    e.persist();
    const { target } = e;
    if (
      Math.ceil(target.scrollTop + target.offsetHeight) === target.scrollHeight &&
      current < pagecount
    ) {
      getValvelist({ current: current + 1 });
    }
  };
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
  const handleUpdate = ()=>{
    const hide=message.loading('正在提交...')
    form
      .validateFields().then(async(values)=>{
      try{
        values.id = id
        const res=await updateAlarm({...values})
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
    form
      .validateFields().then(async(values)=>{
      const hide=message.loading('正在提交...')
      try{
        const res=await addAlarm({...values})
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
        request={(params, sorter, filter) => queryAlarmList({ ...params, sorter, filter })}
        columns={columns}
      />
      <Modal
        title="新建报警"
        visible={createModalVisible}
        onOk={handleAdd}
        onCancel={() => handleModalVisible(false)} 
      >
        <Form form={form} initialValues={formvalues}>
            <Form.Item
              label={
                <span>
                  选择阀 &nbsp;
                  <Tooltip title="选择需要的阀">
                    <QuestionCircleOutlined />
                  </Tooltip>
                </span>
              }
              name="alarmtype_id"
              rules={[
                {
                  required: true,
                  message: '请选择阀!',
                },
              ]}
            >
              <Select
                placeholder="请选择阀..."
                onPopupScroll={handlePopupScrollAlarmtype}
                allowClear
                showSearch
                optionFilterProp="children"
              >
                {Alarmtypelist.data.length &&
                  Alarmtypelist.data.map((obj) => {
                    return <Select.Option value={obj.id} key={obj.id}>{obj.name}</Select.Option>;
                  })}
              </Select>
            </Form.Item>
            <Form.Item
              label={
                <span>
                  选择阀 &nbsp;
                  <Tooltip title="选择需要的阀">
                    <QuestionCircleOutlined />
                  </Tooltip>
                </span>
              }
              name="valve_id"
              rules={[
                {
                  required: true,
                  message: '请选择阀!',
                },
              ]}
            >
              <Select
                placeholder="请选择阀..."
                onPopupScroll={handlePopupScrollValve}
                allowClear
                showSearch
                optionFilterProp="children"
              >
                {Valvelist.data.length &&
                  Valvelist.data.map((obj) => {
                    return <Select.Option value={obj.id} key={obj.id}>{obj.name}</Select.Option>;
                  })}
              </Select>
            </Form.Item>
            <Form.Item
              label={
                <span>
                  选择机台 &nbsp;
                  <Tooltip title="选择需要的机台">
                    <QuestionCircleOutlined />
                  </Tooltip>
                </span>
              }
              name="device_id"
              rules={[
                {
                  required: true,
                  message: '请选择机台!',
                },
              ]}
            >
              <Select
                placeholder="请选择机台..."
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
        </Form>
      </Modal>
      <Modal
        title="编辑报警"
        visible={updateModalVisible}
        onOk={handleUpdate}
        onCancel={()=>{handleUpdateModalVisible(false)}}
      >
        <Form form={form} initialValues={formvalues}>
            <Form.Item
              label={
                <span>
                  选择阀 &nbsp;
                  <Tooltip title="选择需要的阀">
                    <QuestionCircleOutlined />
                  </Tooltip>
                </span>
              }
              name="alarmtype_id"
              rules={[
                {
                  required: true,
                  message: '请选择阀!',
                },
              ]}
            >
              <Select
                placeholder="请选择阀..."
                onPopupScroll={handlePopupScrollAlarmtype}
                allowClear
                showSearch
                optionFilterProp="children"
              >
                {Alarmtypelist.data.length &&
                  Alarmtypelist.data.map((obj) => {
                    return <Select.Option value={obj.id} key={obj.id}>{obj.name}</Select.Option>;
                  })}
              </Select>
            </Form.Item>
            <Form.Item
              label={
                <span>
                  选择阀 &nbsp;
                  <Tooltip title="选择需要的阀">
                    <QuestionCircleOutlined />
                  </Tooltip>
                </span>
              }
              name="valve_id"
              rules={[
                {
                  required: true,
                  message: '请选择阀!',
                },
              ]}
            >
              <Select
                placeholder="请选择阀..."
                onPopupScroll={handlePopupScrollValve}
                allowClear
                showSearch
                optionFilterProp="children"
              >
                {Valvelist.data.length &&
                  Valvelist.data.map((obj) => {
                    return <Select.Option value={obj.id} key={obj.id}>{obj.name}</Select.Option>;
                  })}
              </Select>
            </Form.Item>
            <Form.Item
              label={
                <span>
                  选择机台 &nbsp;
                  <Tooltip title="选择需要的机台">
                    <QuestionCircleOutlined />
                  </Tooltip>
                </span>
              }
              name="device_id"
              rules={[
                {
                  required: true,
                  message: '请选择机台!',
                },
              ]}
            >
              <Select
                placeholder="请选择机台..."
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
        </Form>
      </Modal>
    </PageHeaderWrapper>
  );
};

export default TableList;
