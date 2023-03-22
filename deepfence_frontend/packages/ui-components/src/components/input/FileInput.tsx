import * as LabelPrimitive from '@radix-ui/react-label';
import { cva, VariantProps } from 'cva';
import { ChangeEvent, ComponentProps, forwardRef, useId } from 'react';
import { twMerge } from 'tailwind-merge';

import HelperText from '@/components/input/HelperText';
import { ObjectWithNonNullableValues } from '@/types/utils';

export type SizeType = 'sm' | 'md' | 'lg';

const inputCva = cva(
  [
    'block w-full rounded-lg dark:placeholder-gray-400 cursor-auto',
    'border border-gray-300 dark:border-gray-600',
    'bg-gray-50 dark:bg-gray-700',
    'text-sm text-gray-900 dark:text-white file:text-white file:dark:text-white',
    'disabled:cursor-not-allowed',
    'focus:outline-none ',
    // ring styles
    'focus:ring-1 focus:ring-blue-600',
    'file:border-0 file:cursor-pointer file:hover:bg-gray-700 dark:file:hover:bg-gray-500',
    'file:h-ful file:bg-gray-800 dark:file:bg-gray-600 file:px-4',
  ],
  {
    variants: {
      sizing: {
        sm: 'text-sm file:py-2 file:mr-4',
        md: 'text-sm file:py-3 file:mr-4',
        lg: 'text-base file:py-3.5 file:mr-4',
      },
    },

    defaultVariants: {
      sizing: 'md',
    },
  },
);

export interface TextInputProps
  extends Omit<ComponentProps<'input'>, 'ref' | 'color' | 'size'>,
    ObjectWithNonNullableValues<VariantProps<typeof inputCva>> {
  label?: string;
  helperText?: string;
  required?: boolean;
  accept?: string;
  onChoosen?: (e: ChangeEvent<HTMLInputElement>) => void;
}

export const FileInput = forwardRef<HTMLInputElement, TextInputProps>(
  (
    { sizing, label, helperText, className = '', required, id, onChoosen, ...rest },
    ref,
  ) => {
    const internalId = useId();
    const _id = id ? id : internalId;

    return (
      <div className={twMerge('flex flex-col gap-2 w-full', className)}>
        {label && (
          <LabelPrimitive.Root
            htmlFor={_id}
            className="text-sm font-medium text-gray-900 dark:text-white"
          >
            {required && <span>*</span>}
            {label}
          </LabelPrimitive.Root>
        )}

        <input
          {...rest}
          ref={ref}
          data-testid={`fileinput-${id}`}
          type="file"
          className={twMerge(
            inputCva({
              sizing,
            }),
          )}
          onChange={onChoosen}
        />
        {helperText && <HelperText text={helperText} className="mb-2.5" />}
      </div>
    );
  },
);

export default FileInput;