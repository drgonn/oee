import { DownOutlined, PlusOutlined, QuestionCircleOutlined} from '@ant-design/icons';
import { Button, Divider, Dropdown, Menu, message, Input, Form, Modal, Tooltip, Select, InputNumber ,Upload,DatePicker  } from 'antd';
import React, { useState, useRef, useEffect} from 'react';
import { PageHeaderWrapper } from '@ant-design/pro-layout';
import ProTable, { ProColumns, ActionType } from '@ant-design/pro-table';

import {queryUserList} from "@/pages/user_system/user/service";
import { TableListItem } from './data.d';
import { queryMessageList, updateMessage , addMessage , removeMessage  } from './service';

const { TextArea } = Input;
const handleRemove = async (selectedRows: TableListItem[]) => {
  const hide = message.loading('正在删除');
  if (!selectedRows) return true;
  try {
    const res = await removeMessage({
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
  const [Userlist, setUserlist] = useState({ data: [] });
  const [updateModalVisible, handleUpdateModalVisible] = useState<boolean>(false);
  const [stepFormValues, setStepFormValues] = useState({});
  const [id,setId]=useState(0)
  const [formvalues,setValues]=useState({})
  const [form] = Form.useForm()
  const actionRef = useRef<ActionType>();
  const columns: ProColumns<TableListItem>[] = [
    {
      title: '工作事项',
      dataIndex: 'title',
      valueType: 'text',
      search: true,
    },
    {
      title: '类型',
      dataIndex: 'type',
      valueEnum: {
        1: { text:'提醒'},
        2: { text:'警告'},
      },
      search: true,
    },
    {
      title: '消息内容',
      dataIndex: 'description',
      valueType: 'text',
      search: false,
    },
    {
      title: '消息是否已读',
      dataIndex: 'read',
      valueType: 'text',
      search: true,
    },
    {
      title: '时间',
      dataIndex: 'ct_time',
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

  const getUserlist = async (obj) => {
    const Userdata = await queryUserList(obj);
    if (Userdata.success) {
      const {data} = Userlist
      Userdata.data = data.concat(Userdata.data)
      setUserlist(Userdata);
    }
  };
useEffect(()=>{getUserlist({pageindex:1})},[])
  const handlePopupScrollUser = async (e) => {
    const { current, pagecount } = Userlist;
    e.persist();
    const { target } = e;
    if (
      Math.ceil(target.scrollTop + target.offsetHeight) === target.scrollHeight &&
      current < pagecount
    ) {
      getUserlist({ current: current + 1 });
    }
  };
  const handleUpdate = ()=>{
    const hide=message.loading('正在提交...')
    form
      .validateFields().then(async(values)=>{
      try{
        values.id = id
        const res=await updateMessage({...values})
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
        const res=await addMessage({...values})
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
        request={(params, sorter, filter) => queryMessageList({ ...params, sorter, filter })}
        columns={columns}
      />
      <Modal
        title="新建通知"
        visible={createModalVisible}
        onOk={handleAdd}
        onCancel={() => handleModalVisible(false)} 
      >
        <Form form={form} initialValues={formvalues}>
            <Form.Item
              label={
                <span>
                  选择用户 &nbsp;
                  <Tooltip title="选择需要的用户">
                    <QuestionCircleOutlined />
                  </Tooltip>
                </span>
              }
              name="user_id"
              rules={[
                {
                  required: true,
                  message: '请选择用户!',
                },
              ]}
            >
              <Select
                placeholder="请选择用户..."
                onPopupScroll={handlePopupScrollUser}
                allowClear
                showSearch
                optionFilterProp="children"
              >
                {Userlist.data.length &&
                  Userlist.data.map((obj) => {
                    return <Select.Option value={obj.id} key={obj.id}>{obj.name}</Select.Option>;
                  })}
              </Select>
            </Form.Item>
          <Form.Item
            name='title'
            rules= {[{ required: true, message: '请输入名称!' }]}
            label="通知工作事项"
          >
            <Input placeholder="请输入工作事项" />
          </Form.Item>
          <Form.Item
            name='type'
            rules= {[{ required: true, message: '请输入名称!' }]}
            label="通知类型"
          >
            <Select>
            <Select.Option value={1}>提醒</Select.Option>
            <Select.Option value={2}>警告</Select.Option>
            </Select>
          </Form.Item>
          <Form.Item
            name='description'
            rules= {[{ required: false, message: '请输入名称!' }]}
            label="通知消息内容"
          >
            <TextArea rows={4} />
          </Form.Item>
        </Form>
      </Modal>
      <Modal
        title="编辑通知"
        visible={updateModalVisible}
        onOk={handleUpdate}
        onCancel={()=>{handleUpdateModalVisible(false)}}
      >
        <Form form={form} initialValues={formvalues}>
            <Form.Item
              label={
                <span>
                  选择用户 &nbsp;
                  <Tooltip title="选择需要的用户">
                    <QuestionCircleOutlined />
                  </Tooltip>
                </span>
              }
              name="user_id"
              rules={[
                {
                  required: true,
                  message: '请选择用户!',
                },
              ]}
            >
              <Select
                placeholder="请选择用户..."
                onPopupScroll={handlePopupScrollUser}
                allowClear
                showSearch
                optionFilterProp="children"
              >
                {Userlist.data.length &&
                  Userlist.data.map((obj) => {
                    return <Select.Option value={obj.id} key={obj.id}>{obj.name}</Select.Option>;
                  })}
              </Select>
            </Form.Item>
          <Form.Item
            name='title'
            rules= {[{ required: true, message: '请输入名称!' }]}
            label="通知工作事项"
          >
            <Input placeholder="请输入工作事项" />
          </Form.Item>
          <Form.Item
            name='type'
            rules= {[{ required: true, message: '请输入名称!' }]}
            label="通知类型"
          >
            <Select>
            <Select.Option value={1}>提醒</Select.Option>
            <Select.Option value={2}>警告</Select.Option>
            </Select>
          </Form.Item>
          <Form.Item
            name='description'
            rules= {[{ required: false, message: '请输入名称!' }]}
            label="通知消息内容"
          >
            <TextArea rows={4} />
          </Form.Item>
        </Form>
      </Modal>
    </PageHeaderWrapper>
  );
};

export default TableList;
