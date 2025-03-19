export type Envelope = {
  id: number
  to: string
  from: string
  subject: string
  created_at: string
  animate?: boolean
}

export type Attachment = {
  id: string
  filename: string
}
