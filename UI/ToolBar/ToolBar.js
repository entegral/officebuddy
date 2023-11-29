import {
  View,
  StyleSheet,
} from "react-native";
import { Image } from 'expo-image';
import HamburgerMenuIcon from "../Icons/HamburgerMenuIcon";

export default function ToolBar({
  
}) {
  const src = require('../Images/OfficialLogo.png')
  return (
    <View
      className={'toolbar'}
      style={styles.toolbar}  
    >
      <View
        className={'icon'}
        style={styles.icon}
      >
        <Image
          source={src}
          style={{
            height: 40,
            width: 40,
          }}
        />
      </View>
      <View
        style={styles.MenuContainer}
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

const styles = StyleSheet.create({
  toolbar: {
    display: 'flex',
    flexDirection: 'row',
    alignItems: 'flex-start',
    paddingHorizontal: 16,
    paddingBottom: 8,
    paddingTop: 3,
  },
  icon: {
    height: 40,
    width: 40,
  },
  MenuContainer: {
    paddingVertical: 8,
    marginLeft: 'auto',
  },
});
