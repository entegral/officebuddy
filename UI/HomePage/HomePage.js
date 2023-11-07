import {
  View,
  Text,
} from "react-native";
import ToolBar from "../ToolBar/ToolBar";

export default function HomePage({

}){
  return (
    <View>
      <View
        style={{
          backgroundColor: 'green'
        }}
      >
        <Text
          style={{
            color: '#fff',
          }}
        >
          Home Page
        </Text>
      </View>
    </View>
  )
}