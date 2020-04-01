#version 330 core

in vec3 position;
in vec2 textureCoords;
in vec3 normal;

out VertexData {
    vec2 textureCoords;
    vec3 normal;
} VertexOut;

uniform mat4 transformationMatrix;
uniform mat4 projectionMatrix;
uniform mat4 viewMatrix;

void main()
{
    vec4 worldPos = transformationMatrix * vec4(position, 1.0);
    vec4 relativeToCam = viewMatrix * worldPos;
    gl_Position = projectionMatrix * relativeToCam;

    VertexOut.textureCoords = textureCoords;
    VertexOut.normal = (transformationMatrix * vec4(normal, 0)).xyz;
}
