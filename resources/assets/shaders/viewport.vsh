#version 330 core

in vec3 position;
in vec3 normal;
in vec2 textureCoords;

out vec2 pass_textureCoords;
out vec3 surfaceNormal;
out vec3 toLightVector;
out vec3 toCameraVector;

uniform mat4 transformationMatrix;
uniform mat4 projectionMatrix;
uniform mat4 viewMatrix;
uniform vec3 lightPosition;

void main()
{
    vec4 worldPos = transformationMatrix * vec4(position, 1.0);
    vec4 relativeToCam = viewMatrix * worldPos;
    gl_Position = projectionMatrix * relativeToCam;
    pass_textureCoords = textureCoords;

    surfaceNormal = (transformationMatrix * vec4(normal, 0)).xyz;
    toLightVector = lightPosition - worldPos.xyz;
    toCameraVector = (inverse(viewMatrix) * vec4(0, 0, 0, 1)).xyz - worldPos.xyz;
}
