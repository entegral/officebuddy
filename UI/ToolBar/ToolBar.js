import {
  View,
} from "react-native";
import HamburgerMenuIcon from "../Icons/HamburgerMenuIcon";

export default function ToolBar({
  
}) {
  return (
    <View
      className={'toolbar'}
      style={{
        height: 51,
        flexDirection: 'row',
        paddingHorizontal: 16,
        paddingBottom: 8,
        paddingTop: 2,
      }}  
      // style={{
      //   height: 50,
      //   width: '100%',
      //   flexDirection: 'row',
      //   justifyContent: 'space-between',
      //   alignItems: 'center',
      //   alignSelf: 'flex-start',
      //   paddingHorizontal: 16,
      //   paddingBottom: 8,
      //   paddingTop: 2,
      //   borderColor: '#fff',
      //   borderWidth: 1,
      // }}
    >
      <View
        style={{
          backgroundColor: 'red',
          height: 24,
          width:24,
        }}
      >
        {/* icon goes here */}
      </View>
      <View
      style={{
        backgroundColor: 'blue',
      }}
      >
        <HamburgerMenuIcon
          style={{
            height: 24,
            width: 24,
          }}
          color={'#D2FFAF'}
        />
      </View>
    </View>
  )
};
