import { DownOutlined, PlusOutlined, QuestionCircleOutlined} from '@ant-design/icons';
import { Button, Divider, Dropdown, Menu, message, Input, Form, Modal, Tooltip, Select, InputNumber ,Upload,DatePicker  } from 'antd';
import React, { useState, useRef, useEffect} from 'react';
import { PageHeaderWrapper } from '@ant-design/pro-layout';
import ProTable, { ProColumns, ActionType } from '@ant-design/pro-table';

import { TableListItem } from './data.d';
import { queryAlarmtypeList, updateAlarmtype , addAlarmtype , removeAlarmtype  } from './service';

const { TextArea } = Input;
const handleRemove = async (selectedRows: TableListItem[]) => {
  const hide = message.loading('正在删除');
  if (!selectedRows) return true;
  try {
    const res = await removeAlarmtype({
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
  const [updateModalVisible, handleUpdateModalVisible] = useState<boolean>(false);
  const [stepFormValues, setStepFormValues] = useState({});
  const [id,setId]=useState(0)
  const [formvalues,setValues]=useState({})
  const [form] = Form.useForm()
  const actionRef = useRef<ActionType>();
  const columns: ProColumns<TableListItem>[] = [
    {
      title: '报警代码',
      dataIndex: 'code',
      valueType: 'text',
      search: false,
    },
    {
      title: '含义',
      dataIndex: 'mean',
      valueType: 'text',
      search: false,
    },
    {
      title: '可能原因',
      dataIndex: 'cause',
      valueType: 'text',
      search: false,
    },
    {
      title: '解决方案',
      dataIndex: 'solution',
      valueType: 'text',
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

  const handleUpdate = ()=>{
    const hide=message.loading('正在提交...')
    form
      .validateFields().then(async(values)=>{
      try{
        values.id = id
        const res=await updateAlarmtype({...values})
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
        const res=await addAlarmtype({...values})
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
        request={(params, sorter, filter) => queryAlarmtypeList({ ...params, sorter, filter })}
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
            name='code'
            rules= {[{ required: true, message: '请输入名称!' }]}
            label="报警报警代码"
          >
            <Input placeholder="请输入报警代码" />
          </Form.Item>
          <Form.Item
            name='mean'
            rules= {[{ required: false, message: '请输入名称!' }]}
            label="报警含义"
          >
            <TextArea rows={4} />
          </Form.Item>
          <Form.Item
            name='cause'
            rules= {[{ required: false, message: '请输入名称!' }]}
            label="报警可能原因"
          >
            <TextArea rows={4} />
          </Form.Item>
          <Form.Item
            name='solution'
            rules= {[{ required: false, message: '请输入名称!' }]}
            label="报警解决方案"
          >
            <TextArea rows={4} />
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
            name='code'
            rules= {[{ required: true, message: '请输入名称!' }]}
            label="报警报警代码"
          >
            <Input placeholder="请输入报警代码" />
          </Form.Item>
          <Form.Item
            name='mean'
            rules= {[{ required: false, message: '请输入名称!' }]}
            label="报警含义"
          >
            <TextArea rows={4} />
          </Form.Item>
          <Form.Item
            name='cause'
            rules= {[{ required: false, message: '请输入名称!' }]}
            label="报警可能原因"
          >
            <TextArea rows={4} />
          </Form.Item>
          <Form.Item
            name='solution'
            rules= {[{ required: false, message: '请输入名称!' }]}
            label="报警解决方案"
          >
            <TextArea rows={4} />
          </Form.Item>
        </Form>
      </Modal>
    </PageHeaderWrapper>
  );
};

export default TableList;
