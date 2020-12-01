require 'httparty'
require 'json'

def handler(event:, context:)
    obj = HTTParty.get('https://randomfox.ca/floof/')
    { event: JSON.generate(event), context: JSON.generate(context.inspect), http_request: obj }
end
