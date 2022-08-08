import classnames from 'classnames';
import React, { FC, forwardRef } from 'react'

export interface SignerInputProps extends React.HTMLProps<HTMLInputElement> {
  leftComp?: React.ReactNode
  rightComp?: React.ReactNode
  error?: boolean
}

export const SignerInput = forwardRef<HTMLInputElement, SignerInputProps>((props: SignerInputProps, ref) => {
  const className = classnames(
    'input input-bordered w-full pl-12 pr-14',
    {
      'input-error': props.error,
    }
  )
  return (
    <div className="relative rounded-md shadow-sm">
    {props.leftComp && (
      <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
        {props.leftComp}
      </div>
    )}
    <input
      ref={ref}
      type="text"
      className={className}
      {...props}
    />
    {props.rightComp && (
      <div className="absolute inset-y-0 right-0 flex items-center">
        {props.rightComp}
      </div>
    )}
  </div>
  )
})
