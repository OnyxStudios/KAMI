#version 330 core

in vec2 pass_textureCoords;
in vec3 surfaceNormal;
in vec3 toLightVector;
in vec3 toCameraVector;

out vec4 FragColor;

uniform sampler2D textureSampler;
uniform vec3 lightColor;
uniform vec3 environmentColor;

void main()
{
    vec3 unitNormal = normalize(surfaceNormal);
    vec3 unitLight = normalize(toLightVector);
    vec3 unitCameraVector = normalize(toCameraVector);

    float normalDot = dot(unitNormal, unitLight);
    float brightness = max(normalDot, 0.2);
    vec3 diffuse = brightness * lightColor;
    vec3 lightDirection = -unitLight;
    vec3 reflectedLightDirection = reflect(lightDirection, unitNormal);

    float specularFactor = dot(reflectedLightDirection, unitCameraVector);
    specularFactor = max(specularFactor, 0.0);
    vec3 finalSpecular = specularFactor * lightColor;

    vec4 textureColor = texture(textureSampler, pass_textureCoords);
    if(textureColor.a < 0.5) {
        discard;
    }

    FragColor = vec4(diffuse, 1) * textureColor + vec4(finalSpecular, 1) + vec4(environmentColor, 1);
}
