import { DownOutlined, PlusOutlined, QuestionCircleOutlined} from '@ant-design/icons';
import { Button, Divider, Dropdown, Menu, message, Input, Form, Modal, Tooltip, Select, InputNumber ,Upload,DatePicker  } from 'antd';
import React, { useState, useRef, useEffect} from 'react';
import { PageHeaderWrapper } from '@ant-design/pro-layout';
import ProTable, { ProColumns, ActionType } from '@ant-design/pro-table';

import {queryDeviceList} from "@/pages/oee/device/service";
import { TableListItem } from './data.d';
import { queryWorktimeList, updateWorktime , addWorktime , removeWorktime  } from './service';

const { TextArea } = Input;
const handleRemove = async (selectedRows: TableListItem[]) => {
  const hide = message.loading('正在删除');
  if (!selectedRows) return true;
  try {
    const res = await removeWorktime({
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
          )}
    },
    {
      title: '开始时间',
      dataIndex: 'start_time',
      valueType: 'dateTime',
      search: false,
    },
    {
      title: '结束时间',
      dataIndex: 'end_time',
      valueType: 'dateTime',
      search: false,
    },
    {
      title: '运行时间（秒）',
      dataIndex: 'seconds',
      valueType: 'digit',
      search: false,
    },
    {
      title: '时间类型',
      dataIndex: 'type',
      valueEnum: {
        1: { text:'休息时间'},
        2: { text:'日常管理时间'},
        5: { text:'计划停止时间'},
        6: { text:'日常管理时间'},
        3: { text:'停机时间'},
        4: { text:'运转时间'},
      },
      search: true,
    },
    {
      title: '加工数量',
      dataIndex: 'amount',
      valueType: 'digit',
      search: false,
    },
    {
      title: '良品数量',
      dataIndex: 'good',
      valueType: 'digit',
      search: false,
    },
    {
      title: '使用胶量',
      dataIndex: 'glue',
      valueType: 'digit',
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
        const res=await updateWorktime({...values,...v_time
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
        const res=await addWorktime({...values,...v_time
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
        request={(params, sorter, filter) => queryWorktimeList({ ...params, sorter, filter })}
        columns={columns}
      />
      <Modal
        title="新建机台工作时间"
        visible={createModalVisible}
        onOk={handleAdd}
        onCancel={() => handleModalVisible(false)} 
      >
        <Form form={form} initialValues={formvalues}>
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
          <Form.Item
            name='start_time'
            rules= {[{ required: true, message: '请输入名称!' }]}
            label="机台工作时间开始时间"
          >
            <DatePicker showTime/>
          </Form.Item>
          <Form.Item
            name='end_time'
            rules= {[{ required: true, message: '请输入名称!' }]}
            label="机台工作时间结束时间"
          >
            <DatePicker showTime/>
          </Form.Item>
          <Form.Item
            name='seconds'
            rules= {[{ required: true, message: '请输入名称!' }]}
            label="机台工作时间运行时间（秒）"
          >
            <InputNumber  defaultValue={0}  />
          </Form.Item>
          <Form.Item
            name='type'
            rules= {[{ required: true, message: '请输入名称!' }]}
            label="机台工作时间时间类型"
          >
            <Select>
            <Select.Option value={1}>休息时间</Select.Option>
            <Select.Option value={2}>日常管理时间</Select.Option>
            <Select.Option value={5}>计划停止时间</Select.Option>
            <Select.Option value={6}>日常管理时间</Select.Option>
            <Select.Option value={3}>停机时间</Select.Option>
            <Select.Option value={4}>运转时间</Select.Option>
            </Select>
          </Form.Item>
          <Form.Item
            name='amount'
            rules= {[{ required: false, message: '请输入名称!' }]}
            label="机台工作时间加工数量"
          >
            <InputNumber  defaultValue={0}  />
          </Form.Item>
          <Form.Item
            name='good'
            rules= {[{ required: false, message: '请输入名称!' }]}
            label="机台工作时间良品数量"
          >
            <InputNumber  defaultValue={0}  />
          </Form.Item>
          <Form.Item
            name='glue'
            rules= {[{ required: false, message: '请输入名称!' }]}
            label="机台工作时间使用胶量"
          >
            <Input placeholder="请输入使用胶量" />
          </Form.Item>
        </Form>
      </Modal>
      <Modal
        title="编辑机台工作时间"
        visible={updateModalVisible}
        onOk={handleUpdate}
        onCancel={()=>{handleUpdateModalVisible(false)}}
      >
        <Form form={form} initialValues={formvalues}>
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
          <Form.Item
            name='start_timeu'
            rules= {[{ required: true, message: '请输入名称!' }]}
            label="机台工作时间开始时间"
          >
            <DatePicker showTime/>
          </Form.Item>
          <Form.Item
            name='end_timeu'
            rules= {[{ required: true, message: '请输入名称!' }]}
            label="机台工作时间结束时间"
          >
            <DatePicker showTime/>
          </Form.Item>
          <Form.Item
            name='seconds'
            rules= {[{ required: true, message: '请输入名称!' }]}
            label="机台工作时间运行时间（秒）"
          >
            <InputNumber  defaultValue={0}  />
          </Form.Item>
          <Form.Item
            name='type'
            rules= {[{ required: true, message: '请输入名称!' }]}
            label="机台工作时间时间类型"
          >
            <Select>
            <Select.Option value={1}>休息时间</Select.Option>
            <Select.Option value={2}>日常管理时间</Select.Option>
            <Select.Option value={5}>计划停止时间</Select.Option>
            <Select.Option value={6}>日常管理时间</Select.Option>
            <Select.Option value={3}>停机时间</Select.Option>
            <Select.Option value={4}>运转时间</Select.Option>
            </Select>
          </Form.Item>
          <Form.Item
            name='amount'
            rules= {[{ required: false, message: '请输入名称!' }]}
            label="机台工作时间加工数量"
          >
            <InputNumber  defaultValue={0}  />
          </Form.Item>
          <Form.Item
            name='good'
            rules= {[{ required: false, message: '请输入名称!' }]}
            label="机台工作时间良品数量"
          >
            <InputNumber  defaultValue={0}  />
          </Form.Item>
          <Form.Item
            name='glue'
            rules= {[{ required: false, message: '请输入名称!' }]}
            label="机台工作时间使用胶量"
          >
            <Input placeholder="请输入使用胶量" />
          </Form.Item>
        </Form>
      </Modal>
    </PageHeaderWrapper>
  );
};

export default TableList;
