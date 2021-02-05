import { DownOutlined, PlusOutlined, QuestionCircleOutlined} from '@ant-design/icons';
import { Button, Divider, Dropdown, Menu, message, Input, Form, Modal, Tooltip, Select, InputNumber ,Upload,DatePicker  } from 'antd';
import React, { useState, useRef, useEffect} from 'react';
import { PageHeaderWrapper } from '@ant-design/pro-layout';
import ProTable, { ProColumns, ActionType } from '@ant-design/pro-table';

import {queryProjectList} from "@/pages/oee/project/service";
import {queryUserList} from "@/pages/user_system/user/service";
import { TableListItem } from './data.d';
import { queryPlanList, updatePlan , addPlan , removePlan  } from './service';

const { TextArea } = Input;
const handleRemove = async (selectedRows: TableListItem[]) => {
  const hide = message.loading('正在删除');
  if (!selectedRows) return true;
  try {
    const res = await removePlan({
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
  const [Projectlist, setProjectlist] = useState({ data: [] });
  const [Userlist, setUserlist] = useState({ data: [] });
  const [updateModalVisible, handleUpdateModalVisible] = useState<boolean>(false);
  const [stepFormValues, setStepFormValues] = useState({});
  const [id,setId]=useState(0)
  const [formvalues,setValues]=useState({})
  const [form] = Form.useForm()
  const actionRef = useRef<ActionType>();
  const columns: ProColumns<TableListItem>[] = [
    {
      title: '项目名',
      dataIndex: 'project_name',
      valueType: 'text',
      renderFormItem:()=>{
        return(
            <Form.Item
              label=''
              name="project_id"
            >
              <Select
                placeholder="请选择项目..."
                onPopupScroll={handlePopupScrollProject}
                allowClear
                showSearch
                optionFilterProp="children"
              >
                {Projectlist.data.length &&
                  Projectlist.data.map((obj) => {
                    return <Select.Option value={obj.id} key={obj.id}>{obj.name}</Select.Option>;
                  })}
              </Select>
            </Form.Item>
          )}
    },
    {
      title: '用户名',
      dataIndex: 'user_name',
      valueType: 'text',
      renderFormItem:()=>{
        return(
            <Form.Item
              label=''
              name="user_id"
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
          )}
    },
    {
      title: '工作事项',
      dataIndex: 'name',
      valueType: 'text',
      search: true,
    },
    {
      title: '周',
      dataIndex: 'week',
      valueType: 'digit',
      search: true,
    },
    {
      title: '当前进度及完成情况',
      dataIndex: 'current',
      valueType: 'text',
      search: false,
    },
    {
      title: '后续计划',
      dataIndex: 'follow',
      valueType: 'text',
      search: false,
    },
    {
      title: '提交时间',
      dataIndex: 'time',
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

  const getProjectlist = async (obj) => {
    const Projectdata = await queryProjectList(obj);
    if (Projectdata.success) {
      const {data} = Projectlist
      Projectdata.data = data.concat(Projectdata.data)
      setProjectlist(Projectdata);
    }
  };
useEffect(()=>{getProjectlist({pageindex:1})},[])
  const handlePopupScrollProject = async (e) => {
    const { current, pagecount } = Projectlist;
    e.persist();
    const { target } = e;
    if (
      Math.ceil(target.scrollTop + target.offsetHeight) === target.scrollHeight &&
      current < pagecount
    ) {
      getProjectlist({ current: current + 1 });
    }
  };
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
        const res=await updatePlan({...values})
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
        const res=await addPlan({...values})
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
        request={(params, sorter, filter) => queryPlanList({ ...params, sorter, filter })}
        columns={columns}
      />
      <Modal
        title="新建项目进度"
        visible={createModalVisible}
        onOk={handleAdd}
        onCancel={() => handleModalVisible(false)} 
      >
        <Form form={form} initialValues={formvalues}>
            <Form.Item
              label={
                <span>
                  选择项目 &nbsp;
                  <Tooltip title="选择需要的项目">
                    <QuestionCircleOutlined />
                  </Tooltip>
                </span>
              }
              name="project_id"
              rules={[
                {
                  required: true,
                  message: '请选择项目!',
                },
              ]}
            >
              <Select
                placeholder="请选择项目..."
                onPopupScroll={handlePopupScrollProject}
                allowClear
                showSearch
                optionFilterProp="children"
              >
                {Projectlist.data.length &&
                  Projectlist.data.map((obj) => {
                    return <Select.Option value={obj.id} key={obj.id}>{obj.name}</Select.Option>;
                  })}
              </Select>
            </Form.Item>
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
            name='name'
            rules= {[{ required: true, message: '请输入名称!' }]}
            label="项目进度工作事项"
          >
            <Input placeholder="请输入工作事项" />
          </Form.Item>
          <Form.Item
            name='week'
            rules= {[{ required: true, message: '请输入名称!' }]}
            label="项目进度周"
          >
            <InputNumber  defaultValue={0}  />
          </Form.Item>
          <Form.Item
            name='current'
            rules= {[{ required: false, message: '请输入名称!' }]}
            label="项目进度当前进度及完成情况"
          >
            <TextArea rows={4} />
          </Form.Item>
          <Form.Item
            name='follow'
            rules= {[{ required: false, message: '请输入名称!' }]}
            label="项目进度后续计划"
          >
            <TextArea rows={4} />
          </Form.Item>
        </Form>
      </Modal>
      <Modal
        title="编辑项目进度"
        visible={updateModalVisible}
        onOk={handleUpdate}
        onCancel={()=>{handleUpdateModalVisible(false)}}
      >
        <Form form={form} initialValues={formvalues}>
            <Form.Item
              label={
                <span>
                  选择项目 &nbsp;
                  <Tooltip title="选择需要的项目">
                    <QuestionCircleOutlined />
                  </Tooltip>
                </span>
              }
              name="project_id"
              rules={[
                {
                  required: true,
                  message: '请选择项目!',
                },
              ]}
            >
              <Select
                placeholder="请选择项目..."
                onPopupScroll={handlePopupScrollProject}
                allowClear
                showSearch
                optionFilterProp="children"
              >
                {Projectlist.data.length &&
                  Projectlist.data.map((obj) => {
                    return <Select.Option value={obj.id} key={obj.id}>{obj.name}</Select.Option>;
                  })}
              </Select>
            </Form.Item>
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
            name='name'
            rules= {[{ required: true, message: '请输入名称!' }]}
            label="项目进度工作事项"
          >
            <Input placeholder="请输入工作事项" />
          </Form.Item>
          <Form.Item
            name='week'
            rules= {[{ required: true, message: '请输入名称!' }]}
            label="项目进度周"
          >
            <InputNumber  defaultValue={0}  />
          </Form.Item>
          <Form.Item
            name='current'
            rules= {[{ required: false, message: '请输入名称!' }]}
            label="项目进度当前进度及完成情况"
          >
            <TextArea rows={4} />
          </Form.Item>
          <Form.Item
            name='follow'
            rules= {[{ required: false, message: '请输入名称!' }]}
            label="项目进度后续计划"
          >
            <TextArea rows={4} />
          </Form.Item>
        </Form>
      </Modal>
    </PageHeaderWrapper>
  );
};

export default TableList;
