import { DownOutlined, PlusOutlined, QuestionCircleOutlined} from '@ant-design/icons';
import { Button, Divider, Dropdown, Menu, message, Input, Form, Modal, Tooltip, Select, InputNumber ,Upload } from 'antd';
const { TextArea } = Input;
import React, { useState, useRef, useEffect} from 'react';
import { PageHeaderWrapper } from '@ant-design/pro-layout';
import ProTable, { ProColumns, ActionType } from '@ant-design/pro-table';

import {queryRoleList} from "@/pages/user_system/role/service";
import { TableListItem } from './data.d';
import { queryUserList, updateUser , addUser , removeUser  } from './service';

const handleRemove = async (selectedRows: TableListItem[]) => {
  const hide = message.loading('正在删除');
  if (!selectedRows) return true;
  try {
    const res = await removeUser({
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
  const [Rolelist, setRolelist] = useState({ data: [] });
  const [updateModalVisible, handleUpdateModalVisible] = useState<boolean>(false);
  const [stepFormValues, setStepFormValues] = useState({});
  const [id,setId]=useState(0)
  const [formvalues,setValues]=useState({})
  const [form] = Form.useForm()
  const actionRef = useRef<ActionType>();
  const columns: ProColumns<TableListItem>[] = [
    {
      title: '角色名',
      dataIndex: 'role_name',
      valueType: 'text',
      renderFormItem:()=>{
        return(
            <Form.Item
              label=''
              name="role_id"
            >
              <Select
                placeholder="请选择角色..."
                onPopupScroll={handlePopupScrollRole}
                allowClear
                showSearch
                optionFilterProp="children"
              >
                {Rolelist.data.length &&
                  Rolelist.data.map((obj) => {
                    return <Option value={obj.id}>{obj.name}</Option>;
                  })}
              </Select>
            </Form.Item>
          )}
    },
    {
      title: 'UID',
      dataIndex: 'uid',
      valueType: 'text',
    },
    {
      title: '用户名',
      dataIndex: 'username',
      valueType: 'text',
    },
    {
      title: '手机',
      dataIndex: 'phone',
      valueType: 'text',
    },
    {
      title: '邮箱',
      dataIndex: 'email',
      valueType: 'text',
    },
    {
      title: '邮箱是否绑定',
      dataIndex: 'emailbind',
      valueEnum: {
        1: { text:'已绑定'},
        0: { text:'未绑定'},
      }
    },
    {
      title: '公司',
      dataIndex: 'company',
      valueType: 'text',
    },
    {
      title: '地址',
      dataIndex: 'address',
      valueType: 'text',
    },
    {
      title: '网址',
      dataIndex: 'url',
      valueType: 'text',
    },
    {
      title: '微信昵称',
      dataIndex: 'nickname',
      valueType: 'text',
    },
    {
      title: '微信头像地址',
      dataIndex: 'headimgurl',
      valueType: 'text',
    },
    {
      title: '注册时间',
      dataIndex: 'createDate',
      valueType: 'dateTime',
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

  const getRolelist = async (obj) => {
    const Roledata = await queryRoleList(obj);
    if (Roledata.success) {
      const {data} = Rolelist
      Roledata.data = data.concat(Roledata.data)
      setRolelist(Roledata);
    }
  };
useEffect(()=>{getRolelist({pageindex:1})},[])
  const handlePopupScrollRole = async (e) => {
    const { current, pagecount } = Rolelist;
    e.persist();
    const { target } = e;
    if (
      Math.ceil(target.scrollTop + target.offsetHeight) === target.scrollHeight &&
      current < pagecount
    ) {
      getRolelist({ current: current + 1 });
    }
  };
  const handleUpdate = ()=>{
    const hide=message.loading('正在提交...')
    form
      .validateFields().then(async(values)=>{
      try{
        values.id = id
        const res=await updateUser({...values})
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
        const res=await addUser({...values})
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
        actionRef={actionRef}
        rowKey="id"
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
                  <Menu.Item key="approval">批量审批</Menu.Item>
                </Menu>
              }
            >
              <Button>
                批量操作 <DownOutlined />
              </Button>
            </Dropdown>
          ),
        ]}
        request={(params, sorter, filter) => queryUserList({ ...params, sorter, filter })}
        columns={columns}
      />
      <Modal
        title="新建用户"
        visible={createModalVisible}
        onOk={handleAdd}
        onCancel={() => handleModalVisible(false)} 
      >
        <Form form={form} initialValues={formvalues}>
            <Form.Item
              label={
                <span>
                  选择角色 &nbsp;
                  <Tooltip title="选择需要的角色">
                    <QuestionCircleOutlined />
                  </Tooltip>
                </span>
              }
              name="role_id"
              rules={[
                {
                  required: true,
                  message: '请选择角色!',
                },
              ]}
            >
              <Select
                placeholder="请选择角色..."
                onPopupScroll={handlePopupScrollRole}
                allowClear
                showSearch
                optionFilterProp="children"
              >
                {Rolelist.data.length &&
                  Rolelist.data.map((obj) => {
                    return <Option value={obj.id}>{obj.name}</Option>;
                  })}
              </Select>
            </Form.Item>
          <Form.Item
            name='username'
            rules= {[{ required: false, message: '请输入名称!' }]}
            label="用户用户名"
          >
            <Input placeholder="请输入用户名" />
          </Form.Item>
          <Form.Item
            name='phone'
            rules= {[{ required: false, message: '请输入名称!' }]}
            label="用户手机"
          >
            <Input placeholder="请输入手机" />
          </Form.Item>
          <Form.Item
            name='email'
            rules= {[{ required: false, message: '请输入名称!' }]}
            label="用户邮箱"
          >
            <Input placeholder="请输入邮箱" />
          </Form.Item>
          <Form.Item
            name='emailbind'
            rules= {[{ required: false, message: '请输入名称!' }]}
            label="用户邮箱是否绑定"
          >
            <Select>
            <Option value={1}>已绑定</Option>
            <Option value={0}>未绑定</Option>
            </Select>
          </Form.Item>
          <Form.Item
            name='company'
            rules= {[{ required: false, message: '请输入名称!' }]}
            label="用户公司"
          >
            <Input placeholder="请输入公司" />
          </Form.Item>
          <Form.Item
            name='address'
            rules= {[{ required: false, message: '请输入名称!' }]}
            label="用户地址"
          >
            <Input placeholder="请输入地址" />
          </Form.Item>
          <Form.Item
            name='url'
            rules= {[{ required: false, message: '请输入名称!' }]}
            label="用户网址"
          >
            <Input placeholder="请输入网址" />
          </Form.Item>
          <Form.Item
            name='nickname'
            rules= {[{ required: false, message: '请输入名称!' }]}
            label="用户微信昵称"
          >
            <Input placeholder="请输入微信昵称" />
          </Form.Item>
        </Form>
      </Modal>
      <Modal
        title="编辑用户"
        visible={updateModalVisible}
        onOk={handleUpdate}
        onCancel={()=>{handleUpdateModalVisible(false)}}
      >
        <Form form={form} initialValues={formvalues}>
            <Form.Item
              label={
                <span>
                  选择角色 &nbsp;
                  <Tooltip title="选择需要的角色">
                    <QuestionCircleOutlined />
                  </Tooltip>
                </span>
              }
              name="role_id"
              rules={[
                {
                  required: true,
                  message: '请选择角色!',
                },
              ]}
            >
              <Select
                placeholder="请选择角色..."
                onPopupScroll={handlePopupScrollRole}
                allowClear
                showSearch
                optionFilterProp="children"
              >
                {Rolelist.data.length &&
                  Rolelist.data.map((obj) => {
                    return <Option value={obj.id}>{obj.name}</Option>;
                  })}
              </Select>
            </Form.Item>
          <Form.Item
            name='username'
            rules= {[{ required: false, message: '请输入名称!' }]}
            label="用户用户名"
          >
            <Input placeholder="请输入用户名" />
          </Form.Item>
          <Form.Item
            name='phone'
            rules= {[{ required: false, message: '请输入名称!' }]}
            label="用户手机"
          >
            <Input placeholder="请输入手机" />
          </Form.Item>
          <Form.Item
            name='email'
            rules= {[{ required: false, message: '请输入名称!' }]}
            label="用户邮箱"
          >
            <Input placeholder="请输入邮箱" />
          </Form.Item>
          <Form.Item
            name='emailbind'
            rules= {[{ required: false, message: '请输入名称!' }]}
            label="用户邮箱是否绑定"
          >
            <Select>
            <Option value={1}>已绑定</Option>
            <Option value={0}>未绑定</Option>
            </Select>
          </Form.Item>
          <Form.Item
            name='company'
            rules= {[{ required: false, message: '请输入名称!' }]}
            label="用户公司"
          >
            <Input placeholder="请输入公司" />
          </Form.Item>
          <Form.Item
            name='address'
            rules= {[{ required: false, message: '请输入名称!' }]}
            label="用户地址"
          >
            <Input placeholder="请输入地址" />
          </Form.Item>
          <Form.Item
            name='url'
            rules= {[{ required: false, message: '请输入名称!' }]}
            label="用户网址"
          >
            <Input placeholder="请输入网址" />
          </Form.Item>
          <Form.Item
            name='nickname'
            rules= {[{ required: false, message: '请输入名称!' }]}
            label="用户微信昵称"
          >
            <Input placeholder="请输入微信昵称" />
          </Form.Item>
        </Form>
      </Modal>
    </PageHeaderWrapper>
  );
};

export default TableList;
