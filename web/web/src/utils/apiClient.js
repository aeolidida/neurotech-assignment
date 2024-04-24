import { API_HOST, API_PORT } from '../config';

const BASE_URL = `${API_HOST}:${API_PORT}`;

export const apiClient = async (endpoint, method = 'GET', data = null) => {
  const url = `${BASE_URL}/${endpoint}`;
  const options = {
    method,
  };

  if (data) {
    options.headers = {
        'Content-Type': 'application/json',
      }
    options.body = JSON.stringify(data);
    console.log(options.body)
  }

  try {
    const response = await fetch(url, options);

    let responseData;
    try {
      responseData = await response.json();
    } catch (error) {
      responseData="";
    }

    if (!response.ok) {
        let errText = responseData['error'];
        throw new Error(`API error: ${response.status} ${errText}`);
    }

    return responseData;
  } catch (error) {
    console.error('API error:', error);
    throw error;
  }
};