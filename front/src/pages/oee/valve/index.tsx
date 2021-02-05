import { DownOutlined, PlusOutlined, QuestionCircleOutlined} from '@ant-design/icons';
import { Button, Divider, Dropdown, Menu, message, Input, Form, Modal, Tooltip, Select, InputNumber ,Upload,DatePicker  } from 'antd';
import React, { useState, useRef, useEffect} from 'react';
import { PageHeaderWrapper } from '@ant-design/pro-layout';
import ProTable, { ProColumns, ActionType } from '@ant-design/pro-table';

import {queryValvetypeList} from "@/pages/oee/valvetype/service";
import {queryDeviceList} from "@/pages/oee/device/service";
import { TableListItem } from './data.d';
import { queryValveList, updateValve , addValve , removeValve  } from './service';

const { TextArea } = Input;
const handleRemove = async (selectedRows: TableListItem[]) => {
  const hide = message.loading('正在删除');
  if (!selectedRows) return true;
  try {
    const res = await removeValve({
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
  const [Valvetypelist, setValvetypelist] = useState({ data: [] });
  const [Devicelist, setDevicelist] = useState({ data: [] });
  const [updateModalVisible, handleUpdateModalVisible] = useState<boolean>(false);
  const [stepFormValues, setStepFormValues] = useState({});
  const [id,setId]=useState(0)
  const [formvalues,setValues]=useState({})
  const [form] = Form.useForm()
  const actionRef = useRef<ActionType>();
  const columns: ProColumns<TableListItem>[] = [
    {
      title: '类型名',
      dataIndex: 'valvetype_name',
      valueType: 'text',
      renderFormItem:()=>{
        return(
            <Form.Item
              label=''
              name="valvetype_id"
            >
              <Select
                placeholder="请选择类型..."
                onPopupScroll={handlePopupScrollValvetype}
                allowClear
                showSearch
                optionFilterProp="children"
              >
                {Valvetypelist.data.length &&
                  Valvetypelist.data.map((obj) => {
                    return <Select.Option value={obj.id} key={obj.id}>{obj.name}</Select.Option>;
                  })}
              </Select>
            </Form.Item>
          )}
    },
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
      title: '编号',
      dataIndex: 'sn',
      valueType: 'text',
      search: true,
    },
    {
      title: '名称',
      dataIndex: 'name',
      valueType: 'text',
      search: true,
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

  const getValvetypelist = async (obj) => {
    const Valvetypedata = await queryValvetypeList(obj);
    if (Valvetypedata.success) {
      const {data} = Valvetypelist
      Valvetypedata.data = data.concat(Valvetypedata.data)
      setValvetypelist(Valvetypedata);
    }
  };
useEffect(()=>{getValvetypelist({pageindex:1})},[])
  const handlePopupScrollValvetype = async (e) => {
    const { current, pagecount } = Valvetypelist;
    e.persist();
    const { target } = e;
    if (
      Math.ceil(target.scrollTop + target.offsetHeight) === target.scrollHeight &&
      current < pagecount
    ) {
      getValvetypelist({ current: current + 1 });
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
        const res=await updateValve({...values})
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
        const res=await addValve({...values})
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
        request={(params, sorter, filter) => queryValveList({ ...params, sorter, filter })}
        columns={columns}
      />
      <Modal
        title="新建点胶阀"
        visible={createModalVisible}
        onOk={handleAdd}
        onCancel={() => handleModalVisible(false)} 
      >
        <Form form={form} initialValues={formvalues}>
            <Form.Item
              label={
                <span>
                  选择类型 &nbsp;
                  <Tooltip title="选择需要的类型">
                    <QuestionCircleOutlined />
                  </Tooltip>
                </span>
              }
              name="valvetype_id"
              rules={[
                {
                  required: true,
                  message: '请选择类型!',
                },
              ]}
            >
              <Select
                placeholder="请选择类型..."
                onPopupScroll={handlePopupScrollValvetype}
                allowClear
                showSearch
                optionFilterProp="children"
              >
                {Valvetypelist.data.length &&
                  Valvetypelist.data.map((obj) => {
                    return <Select.Option value={obj.id} key={obj.id}>{obj.name}</Select.Option>;
                  })}
              </Select>
            </Form.Item>
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
            name='sn'
            rules= {[{ required: true, message: '请输入名称!' }]}
            label="点胶阀编号"
          >
            <Input placeholder="请输入编号" />
          </Form.Item>
          <Form.Item
            name='name'
            rules= {[{ required: true, message: '请输入名称!' }]}
            label="点胶阀名称"
          >
            <Input placeholder="请输入名称" />
          </Form.Item>
        </Form>
      </Modal>
      <Modal
        title="编辑点胶阀"
        visible={updateModalVisible}
        onOk={handleUpdate}
        onCancel={()=>{handleUpdateModalVisible(false)}}
      >
        <Form form={form} initialValues={formvalues}>
            <Form.Item
              label={
                <span>
                  选择类型 &nbsp;
                  <Tooltip title="选择需要的类型">
                    <QuestionCircleOutlined />
                  </Tooltip>
                </span>
              }
              name="valvetype_id"
              rules={[
                {
                  required: true,
                  message: '请选择类型!',
                },
              ]}
            >
              <Select
                placeholder="请选择类型..."
                onPopupScroll={handlePopupScrollValvetype}
                allowClear
                showSearch
                optionFilterProp="children"
              >
                {Valvetypelist.data.length &&
                  Valvetypelist.data.map((obj) => {
                    return <Select.Option value={obj.id} key={obj.id}>{obj.name}</Select.Option>;
                  })}
              </Select>
            </Form.Item>
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
            name='sn'
            rules= {[{ required: true, message: '请输入名称!' }]}
            label="点胶阀编号"
          >
            <Input placeholder="请输入编号" />
          </Form.Item>
          <Form.Item
            name='name'
            rules= {[{ required: true, message: '请输入名称!' }]}
            label="点胶阀名称"
          >
            <Input placeholder="请输入名称" />
          </Form.Item>
        </Form>
      </Modal>
    </PageHeaderWrapper>
  );
};

export default TableList;
