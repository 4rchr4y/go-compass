import { twMerge } from 'tailwind-merge'

interface IssuesIconProps {
  className?: string
}

export const IssuesIcon = ({ className }: IssuesIconProps) => {
  return (
    <svg
      className={twMerge('flex items-center fill-current', className)}
      viewBox="0 0 16 16"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path
        d="M15.0119 12.8L15.005 12.8L3.38807 12.8L3.35002 12.8L3.35001 12.845L3.35001 13.565C3.35001 13.9542 3.35053 14.2055 3.36616 14.3967C3.38114 14.5801 3.40657 14.6481 3.42358 14.6814C3.48829 14.8085 3.59155 14.9117 3.71856 14.9764C3.75194 14.9934 3.81992 15.0189 4.00329 15.0338C4.19455 15.0495 4.44583 15.05 4.835 15.05L13.565 15.05C13.9542 15.05 14.2055 15.0495 14.3967 15.0338C14.5801 15.0189 14.6481 14.9934 14.6814 14.9764C14.8085 14.9117 14.9117 14.8085 14.9764 14.6814C14.9934 14.6481 15.0189 14.5801 15.0339 14.3967C15.0495 14.2055 15.05 13.9542 15.05 13.565V12.845L15.05 12.8L15.0119 12.8ZM13.7 11.45L13.7 7.715C13.7 6.82178 13.6995 6.19246 13.6593 5.70105C13.6198 5.21753 13.5454 4.92727 13.4303 4.70137C13.193 4.23567 12.8143 3.85705 12.3486 3.61976C12.1227 3.50466 11.8325 3.43018 11.349 3.39067C10.8576 3.35053 10.2282 3.35 9.33501 3.35H9.06501C8.17179 3.35 7.54247 3.35053 7.05106 3.39067C6.56754 3.43018 6.27728 3.50466 6.05138 3.61976C5.58568 3.85705 5.20706 4.23567 4.96977 4.70137C4.85467 4.92727 4.78019 5.21753 4.74068 5.70105C4.70053 6.19246 4.70001 6.82178 4.70001 7.715L4.70001 11.45L13.7 11.45ZM15.05 11.45V7.68565C15.05 6.82845 15.05 6.14396 15.0049 5.59112C14.9585 5.02411 14.8613 4.53641 14.6331 4.08849C14.2664 3.36877 13.6812 2.78362 12.9615 2.4169C12.5136 2.18867 12.0259 2.09148 11.4589 2.04516C10.9061 1.99999 10.2216 1.99999 9.36438 2H9.03564C8.17845 1.99999 7.49396 1.99999 6.94113 2.04516C6.37412 2.09148 5.88642 2.18867 5.4385 2.4169C4.71878 2.78362 4.13363 3.36877 3.76691 4.08849C3.53868 4.53641 3.44149 5.02411 3.39517 5.59112C3.35 6.14395 3.35 6.82844 3.35001 7.68563L3.35001 11.45C3.24772 11.45 3.13933 11.4506 3.04448 11.4584C2.9277 11.4679 2.77341 11.4915 2.61427 11.5726C2.40258 11.6805 2.23048 11.8526 2.12262 12.0643C2.04153 12.2234 2.01792 12.3777 2.00838 12.4945C1.99995 12.5976 1.99998 12.7168 2 12.8267L2 12.8268L2.00001 13.591C1.99999 13.9468 1.99998 14.2538 2.02064 14.5066C2.04244 14.7735 2.09058 15.0389 2.22072 15.2943C2.41486 15.6754 2.72465 15.9851 3.10567 16.1793C3.36108 16.3094 3.6265 16.3576 3.89336 16.3794C4.1462 16.4 4.45316 16.4 4.80899 16.4L13.591 16.4C13.9468 16.4 14.2538 16.4 14.5066 16.3794C14.7735 16.3576 15.0389 16.3094 15.2943 16.1793C15.6754 15.9851 15.9852 15.6754 16.1793 15.2943C16.3094 15.0389 16.3576 14.7735 16.3794 14.5066C16.4 14.2538 16.4 13.9469 16.4 13.591V13.591L16.4 12.845L16.4 12.8268C16.4 12.7169 16.4001 12.5977 16.3916 12.4945C16.3821 12.3777 16.3585 12.2234 16.2774 12.0643C16.1695 11.8526 15.9974 11.6805 15.7857 11.5726C15.6266 11.4915 15.4723 11.4679 15.3555 11.4584C15.2607 11.4506 15.1523 11.45 15.05 11.45ZM7.59717 7.88451H9.01212L8.27217 9.79398C8.16547 10.0694 8.45542 10.2127 8.64331 9.99198L10.9304 7.26774C10.9768 7.21312 11 7.1585 11 7.09932C11 6.99463 10.9165 6.91725 10.8028 6.91725H9.3879L10.1255 5.0055C10.2322 4.73239 9.94228 4.58674 9.75671 4.8075L7.46728 7.53174C7.42089 7.58864 7.40001 7.64099 7.40001 7.70016C7.40001 7.80713 7.48351 7.88451 7.59717 7.88451Z"
        fill-rule="evenodd"
        clip-rule="evenodd"
      />
    </svg>
  )
}

export default IssuesIcon
