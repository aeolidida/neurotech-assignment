using System;
using System.Text.Json;
using System.Text.Json.Serialization;

namespace Desktop.Models
{
    public class Patient
    {
        [JsonPropertyName("fullName")]
        public string FullName { get; set; } = string.Empty;

        [JsonPropertyName("birthday")]
        [JsonConverter(typeof(JsonDateConverter))]
        public DateTime Birthday { get; set; } = DateTime.MinValue;

        [JsonPropertyName("gender")]
        [JsonConverter(typeof(JsonGenderConverter))]
        public Gender Gender { get; set; } = Gender.Unknown;

        [JsonPropertyName("guid")]
        public Guid GUID { get; set; } = Guid.Empty;

        public Patient Clone()
        {
            return new Patient
            {
                FullName = this.FullName,
                Birthday = this.Birthday,
                Gender = this.Gender,
                GUID = this.GUID
            };
        }
    }

    public enum Gender
    {
        Male,
        Female,
        Unknown,
    }

    public class JsonDateConverter : JsonConverter<DateTime>
    {
        public override DateTime Read(ref Utf8JsonReader reader, Type typeToConvert, JsonSerializerOptions options)
        {
            var dateString = reader.GetString();
            if (dateString != null)
            {
                return DateTime.ParseExact(dateString, "yyyy-MM-dd", null);
            }
            throw new JsonException("Invalid date format");
        }

        public override void Write(Utf8JsonWriter writer, DateTime value, JsonSerializerOptions options)
        {
            writer.WriteStringValue(value.ToString("yyyy-MM-dd"));
        }
    }

    public class JsonGenderConverter : JsonConverter<Gender>
    {
        public override Gender Read(ref Utf8JsonReader reader, Type typeToConvert, JsonSerializerOptions options)
        {
            if (reader.TokenType == JsonTokenType.Number)
            {
                var value = reader.GetInt32();
                return (Gender)value;
            }
            throw new JsonException("Invalid gender format");
        }

        public override void Write(Utf8JsonWriter writer, Gender value, JsonSerializerOptions options)
        {
            writer.WriteNumberValue((int)value);
        }
    }
}