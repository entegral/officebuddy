import { TouchableOpacity, Text, View } from 'react-native';

export default function LogInType({
  headerBase,
  buttonBase,
  buttonText
}) {
  return(
    <View>
      <View>
        <Text style={headerBase}>Workmates</Text>
      </View>
      <TouchableOpacity
        style={buttonBase}
      >
        <Text style={buttonText}>
          Continue With email
        </Text>
      </TouchableOpacity>
      <TouchableOpacity
        style={buttonBase}
      >
        <Text style={buttonText}>
          Continue With google
        </Text>
      </TouchableOpacity>
    </View>
  )
}
