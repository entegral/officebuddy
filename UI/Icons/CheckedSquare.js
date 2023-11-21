import Svg, { Path, G, Defs, Rect, ClipPath } from 'react-native-svg'
import { View } from 'react-native';

function CheckedSquareIcon({
  onClick,
  color = 'white',
  height = 16,
  width = 16,
  }) {
  let clickProps;

  if (onClick) {
    clickProps = { onClick, role: 'button', tabIndex: '0' };
  }
  return (
    <View style={{
      height: height,
      width: width,
    }}
    {...clickProps}
    >
      <Svg
        className={'icon'}
        width='100%' 
        height='100%' 
        viewBox='0 0 16 16' 
        fill='none'
        xmlns='http://www.w3.org/2000/svg'
      >
        <G clip-path='url(#clip0_910_446)'>
          <Path d='M12.0001 0.571411H4.0001C2.10656 0.571411 0.571533 2.10644 0.571533 3.99998V12C0.571533 13.8936 2.10656 15.4286 4.0001 15.4286H12.0001C13.8937 15.4286 15.4287 13.8936 15.4287 12V3.99998C15.4287 2.10644 13.8937 0.571411 12.0001 0.571411Z' stroke={color} strokeWidth='1.71429' strokeLinecap='round' strokeLinejoin='round'/>
          <Path d='M11.3339 5.42859L6.76252 11.1429L4.47681 9.42859' stroke={color} strokeWidth='1.71429' strokeLinecap='round' strokeLinejoin='round'/>
        </G>
        <Defs>
          <ClipPath id='clip0_910_446'>
            <Rect width='16' height='16' fill={color}/>
          </ClipPath>
        </Defs>
      </Svg>
    </View>
  )
};
export default CheckedSquareIcon;
