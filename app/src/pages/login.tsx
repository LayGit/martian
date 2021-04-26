import { Form, Input, Button } from 'antd'
import styles from './login.less'

const formLayout = {
  labelCol: { span: 8 },
  wrapperCol: { span: 16 },
}

const tailLayout = {
  wrapperCol: { offset: 8, span: 16 },
}

export default function LoginPage() {
  return (
    <div className={styles.container}>
      <div className={styles.content}>
        <div className={styles.top}>
          <div className={styles.header}><span>Martian</span></div>
          <div className={styles.desc}>What's inside? Martian and You knows!</div>
        </div>
        <div className={styles.main}>
          <div className={styles.form}>
            <LoginStep1 />
          </div>
        </div>
      </div>
    </div>
  )
}

export function LoginStep1() {
  const fetchRepo = (args: any) => {
    
  }

  return (
    <Form {...formLayout} name="login" onFinish={fetchRepo}>
      <Form.Item
        label="Github User"
        name="gitUser"
        rules={[{ required: true, message: 'Please input your github username!' }]}
      >
        <Input />
      </Form.Item>

      <Form.Item {...tailLayout}>
        <Button type="primary" htmlType="submit">Next</Button>
      </Form.Item>
    </Form>
  )
}