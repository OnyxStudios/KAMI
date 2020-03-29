#version 330 core

in vec2 pass_textureCoords;
in vec3 surfaceNormal;
//in vec3 toLightVector;
//in vec3 toCameraVector;

out vec4 FragColor;

uniform sampler2D textureSampler;
//uniform vec3 lightColor;

void main()
{
    vec4 textureColor = texture(textureSampler, pass_textureCoords);
    if(textureColor.a < 0.5) {
        discard;
    }

    FragColor = textureColor;
}
