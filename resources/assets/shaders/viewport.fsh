#version 330 core

in VertexData {
    vec2 textureCoords;
    vec3 normal;
} VertexIn;

out vec4 FragColor;

uniform sampler2D textureSampler;

void main()
{
    vec4 textureColor = texture(textureSampler, VertexIn.textureCoords);
    if(textureColor.a < 0.5) {
        discard;
    }

    FragColor = textureColor;
}
