import React, { FC } from 'react'
import ClipLoader from 'react-spinners/ClipLoader'

export interface LoadingSpinnerProps {
  
}

export const LoadingSpinner: FC<LoadingSpinnerProps> = (props) => {
  return (
    <div className="w-screen h-screen flex items-center justify-center">
      <ClipLoader size={32} color="#999999" />
    </div>
  )
}
