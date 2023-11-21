import { View } from "react-native";
import Svg, { Path } from 'react-native-svg';

function HamburgerMenuIcon({onClick, color = 'white'}) {
  let clickProps;

  if (onClick) {
    clickProps = { onClick, role: 'button', tabIndex: '0' };
  }
//https://docs.expo.dev/versions/latest/sdk/svg/
  return (
    <Svg
      className={'icon'}
      {...clickProps}
      data-testid='CeleritasHamburgerMenuIcon'
      width='24'
      height='25'
      viewBox='0 0 24 25'
      fill='none'
    >
      <Path
        d='M2 6.00659C2 5.45431 2.44772 5.00659 3 5.00659H21C21.5523 5.00659 22 5.45431 22 6.00659C22 6.55888 21.5523 7.00659 21 7.00659H3C2.44772 7.00659 2 6.55888 2 6.00659Z'
        fill={color}
      ></Path>
      <Path
        d='M2 12.0066C2 11.4543 2.44772 11.0066 3 11.0066H21C21.5523 11.0066 22 11.4543 22 12.0066C22 12.5589 21.5523 13.0066 21 13.0066H3C2.44772 13.0066 2 12.5589 2 12.0066Z'
        fill={color}
      ></Path>
      <Path
        d='M3 17.0066C2.44772 17.0066 2 17.4543 2 18.0066C2 18.5589 2.44772 19.0066 3 19.0066H21C21.5523 19.0066 22 18.5589 22 18.0066C22 17.4543 21.5523 17.0066 21 17.0066H3Z'
        fill={color}
      ></Path>
    </Svg>
  );
};

export default HamburgerMenuIcon;
