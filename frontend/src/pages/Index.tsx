import React, { FC } from 'react'
import { useAuth } from '../hooks/useAuth';

export interface IndexProps {}

export const Index: FC<IndexProps> = () => {
  const auth = useAuth()
  return (
    <div>
      <button onClick={() => auth()}>
        Authenticate
      </button>
    </div>
  )
}
