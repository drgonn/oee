import { DownOutlined, PlusOutlined, QuestionCircleOutlined} from '@ant-design/icons';
import { Button, Divider, Dropdown, Menu, message, Input, Form, Modal, Tooltip, Select, InputNumber ,Upload,DatePicker  } from 'antd';
import React, { useState, useRef, useEffect} from 'react';
import { PageHeaderWrapper } from '@ant-design/pro-layout';
import ProTable, { ProColumns, ActionType } from '@ant-design/pro-table';

import { getToken } from '@/utils/authority';
import { UploadOutlined } from '@ant-design/icons';
import { TableListItem } from './data.d';
import { queryDeviceList, updateDevice , addDevice , removeDevice  } from './service';

const { TextArea } = Input;
const handleRemove = async (selectedRows: TableListItem[]) => {
  const hide = message.loading('正在删除');
  if (!selectedRows) return true;
  try {
    const res = await removeDevice({
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
  const [filename,setFilename]=useState('')
  const [loading,setLoading]=useState(false)
  const columns: ProColumns<TableListItem>[] = [
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
      title: 'ip地址',
      dataIndex: 'ip',
      valueType: 'text',
      search: false,
    },
    {
      title: '图片',
      dataIndex: 'img',
      valueType: 'text',
      search: false,
    },
    {
      title: '类型',
      dataIndex: 'type',
      valueEnum: {
        1: { text:'三轴'},
        2: { text:'五轴'},
      },
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

  const handleUpdate = ()=>{
    const hide=message.loading('正在提交...')
    form
      .validateFields().then(async(values)=>{
      try{
        values.id = id
        const res=await updateDevice({...values})
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
        const res=await addDevice({...values,file_name:filename})
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
  const props = {
    name: 'file',
    // accept:'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet, application/vnd.ms-excel',
    listType:"text",
    action:`/api/upload?token=${getToken()}`,
    showUploadList:false,
    beforeUpload(file){
      setLoading(true)
      const checkType=()=>{
        const nameArr=file.name.split('.')
        const nameType=nameArr[nameArr.length-1]
        // if(nameType==='xls'|| nameType==='xlsx' ){
        //   return true
        // }
        return true
      }

      if(!checkType()){
        setLoading(false)
        message.error('文件格式错误，仅支持EXCEL格式上传！')
      }
      const isLt10M = file.size / 1024 / 1024 < 10;

      if (!isLt10M) {
        setLoading(false)
        message.error('文件大小为10M以内!');
      }

      return checkType() && isLt10M ;
    },
    onChange(info) {
      if (info.file.status === 'done') {

        setLoading(false)


        if(info.file.response.ret===false){
          message.error(info.file.response.errmsg ||'上传失败')
          return
        }


        setFilename(info.file.response.upload_file)

        message.success('文件上传成功！')

      }}
  };
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
        request={(params, sorter, filter) => queryDeviceList({ ...params, sorter, filter })}
        columns={columns}
      />
      <Modal
        title="新建机器"
        visible={createModalVisible}
        onOk={handleAdd}
        onCancel={() => handleModalVisible(false)} 
      >
        <Form form={form} initialValues={formvalues}>
          <Form.Item
            name='sn'
            rules= {[{ required: true, message: '请输入名称!' }]}
            label="机器编号"
          >
            <Input placeholder="请输入编号" />
          </Form.Item>
          <Form.Item
            name='name'
            rules= {[{ required: true, message: '请输入名称!' }]}
            label="机器名称"
          >
            <Input placeholder="请输入名称" />
          </Form.Item>
          <Form.Item
            name='ip'
            rules= {[{ required: true, message: '请输入名称!' }]}
            label="机器ip地址"
          >
            <Input placeholder="请输入ip地址" />
          </Form.Item>
          <Form.Item
            name='img'
            rules= {[{ required: true, message: '请输入名称!' }]}
            label="机器图片"
          >
            <Input placeholder="请输入图片" />
          </Form.Item>
          <Form.Item
            name='type'
            rules= {[{ required: true, message: '请输入名称!' }]}
            label="机器类型"
          >
            <Select>
            <Select.Option value={1}>三轴</Select.Option>
            <Select.Option value={2}>五轴</Select.Option>
            </Select>
          </Form.Item>
          <Form.Item
            label="上传文件"
            name="file"
            valuePropName="file"
            rules={[
              {
                required: true,
                message: '请上传文件!',
              },
            ]}
          >
            <Upload {...props}>
              <Button loading={loading}>
                {
                  loading?"正在上传":(<><UploadOutlined /> 点击上传</>)
                }

              </Button>
            </Upload>

          </Form.Item>
        </Form>
      </Modal>
      <Modal
        title="编辑机器"
        visible={updateModalVisible}
        onOk={handleUpdate}
        onCancel={()=>{handleUpdateModalVisible(false)}}
      >
        <Form form={form} initialValues={formvalues}>
          <Form.Item
            name='sn'
            rules= {[{ required: true, message: '请输入名称!' }]}
            label="机器编号"
          >
            <Input placeholder="请输入编号" />
          </Form.Item>
          <Form.Item
            name='name'
            rules= {[{ required: true, message: '请输入名称!' }]}
            label="机器名称"
          >
            <Input placeholder="请输入名称" />
          </Form.Item>
          <Form.Item
            name='ip'
            rules= {[{ required: true, message: '请输入名称!' }]}
            label="机器ip地址"
          >
            <Input placeholder="请输入ip地址" />
          </Form.Item>
          <Form.Item
            name='img'
            rules= {[{ required: true, message: '请输入名称!' }]}
            label="机器图片"
          >
            <Input placeholder="请输入图片" />
          </Form.Item>
          <Form.Item
            name='type'
            rules= {[{ required: true, message: '请输入名称!' }]}
            label="机器类型"
          >
            <Select>
            <Select.Option value={1}>三轴</Select.Option>
            <Select.Option value={2}>五轴</Select.Option>
            </Select>
          </Form.Item>
        </Form>
      </Modal>
    </PageHeaderWrapper>
  );
};

export default TableList;
